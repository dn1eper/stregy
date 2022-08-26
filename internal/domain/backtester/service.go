package backtester

import (
	"context"
	"encoding/binary"
	"os"
	"stregy/internal/domain/backtester/exchange"
	"stregy/internal/domain/broker"
	"stregy/internal/domain/dataseries"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	"stregy/internal/domain/strategy/core"
	"stregy/internal/domain/tick"
	"stregy/internal/domain/tradingac"
	userstrategy "stregy/user/strategy"

	log "github.com/sirupsen/logrus"
)

type Service interface {
	Start(ctx context.Context, b *Backtester) error
	Run(ctx context.Context, b *Backtester) error
	Create(ctx context.Context, dto BacktesterDTO) (*Backtester, error)
	Get(id string) (*Backtester, error)
}

type service struct {
	repository      Repository
	tickService     tick.Service
	quoteService    quote.Service
	exgAccService   exgaccount.Service
	strategyService strategy.Service
	positionService position.Service
	executor        Executor
}

func NewService(
	repository Repository,
	tickService tick.Service,
	quoteService quote.Service,
	exgAccService exgaccount.Service,
	positionService position.Service,
	strategyService strategy.Service,
	executor Executor,
) Service {
	return &service{
		repository:      repository,
		tickService:     tickService,
		quoteService:    quoteService,
		exgAccService:   exgAccService,
		strategyService: strategyService,
		positionService: positionService,
		executor:        executor,
	}
}

func (s *service) Create(ctx context.Context, dto BacktesterDTO) (*Backtester, error) {
	strat := strategy.Strategy{ID: dto.StrategyID}
	bt := Backtester{
		Strategy:            strat,
		StartDate:           dto.StartDate,
		EndDate:             dto.EndDate,
		Symbol:              dto.Symbol,
		Timeframe:           dto.Timeframe,
		HighOrderResolution: dto.HighOrderResolution,
		Status:              Created,
	}
	return s.repository.CreateBacktest(ctx, bt)
}

func (s *service) Start(ctx context.Context, bt *Backtester) error {
	return s.executor.Execute(ctx, bt)
}

func (s *service) Get(id string) (*Backtester, error) {
	return s.repository.GetBacktest(id)
}

// AR is used instead of ATR for simplicity
// Order resolution: all orders except market orders are executed with order price
func (s *service) Run(ctx context.Context, bt *Backtester) error {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	// _ = userstrategy.NewStrategy().(*userstrategy.Strategy)
	impl := userstrategy.NewStrategy()
	bt.Strategy.Implementation = impl
	stratConf := impl.Config()
	tradingAccount := tradingac.NewAccount()
	dataSeries := dataseries.NewDataSeries(stratConf.BarsNeeded)
	broker := broker.NewBroker()

	quoteFeed := s.quoteService.Get(ctx, bt.Symbol, bt.StartDate, bt.EndDate, bt.Timeframe)
	var tickFeed <-chan tick.Tick
	if bt.HighOrderResolution {
		tickFeed = s.tickService.GetHistorical(ctx, bt.Symbol, bt.StartDate, bt.EndDate)
	} else {
		quoteFeed, tickFeed = ticksFromQuotes(quoteFeed, bt.Timeframe)
	}

	SaveTicksToBinary(tickFeed)

	// final objects
	barsNeeded := stratConf.BarsNeeded
	if stratConf.ATRperiod > barsNeeded {
		barsNeeded = stratConf.ATRperiod
	}
	exchange := exchange.NewExchange(broker, tickFeed, quoteFeed, bt.Timeframe, barsNeeded)
	broker.Configure(&dataSeries, tradingAccount, &exchange, impl)
	core.Configure(broker, &dataSeries, tradingAccount)

	return exchange.Run()
}

type Tick struct {
	sec     int64
	nanosec int32
	price   float64
}

func SaveTicksToBinary(tickFeed <-chan tick.Tick) {
	ticks := make([]Tick, 0)
	for tk := range tickFeed {
		// fmt.Printf("%s %s %f\n", tick.Time, tick.Symbol, tick.Price)
		ticks = append(ticks, Tick{tk.Time.Unix(), int32(tk.Time.Nanosecond()), tk.Price})
	}

	f, err := os.Create("stregy/tests/backtester/repository/ticks.bin")
	if err != nil {
		log.Error(err)
		return
	}
	defer f.Close()
	err = binary.Write(f, binary.LittleEndian, ticks)
	if err != nil {
		log.Error(err)
		return
	}
}
