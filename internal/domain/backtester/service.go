package backtester

import (
	"context"
	"fmt"
	"math"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	"stregy/internal/domain/tick"
	"time"
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

var Quotes []quote.Quote
var ATR float64

var activeOrders []order.Order
var activePositions []position.Position

func ActiveOrders() []order.Order {
	activeOrdersCopy := make([]order.Order, len(activeOrders))
	copy(activeOrdersCopy, activeOrders)
	return activeOrdersCopy
}
func ActivePositions() []position.Position {
	activePositionsCopy := make([]position.Position, len(activePositions))
	copy(activePositionsCopy, activePositions)
	return activePositionsCopy
}

func Quote(i int) quote.Quote {
	return Quotes[len(Quotes)-1-i]
}
func Open(i int) float64 {
	return Quotes[len(Quotes)-1-i].Open
}
func High(i int) float64 {
	return Quotes[len(Quotes)-1-i].High
}
func Low(i int) float64 {
	return Quotes[len(Quotes)-1-i].Low
}
func Close(i int) float64 {
	return Quotes[len(Quotes)-1-i].Close
}
func Volume(i int) float64 {
	return Quotes[len(Quotes)-1-i].Volume
}

func (s *service) Start(ctx context.Context, bt *Backtester) error {
	return s.executor.Execute(ctx, bt)
}

// AR is used instead of ATR for simplicity
// Order resolution: all orders except market orders are executed with order price
func (s *service) Run(ctx context.Context, bt *Backtester) error {
	// tfDuration := time.Minute * time.Duration(bt.Timeframe)

	// var orders []order.Order
	// var positions []position.Position

	// quoteChan := s.quoteService.Get(ctx, bt.Symbol, bt.StartDate, bt.EndDate, bt.Timeframe)
	// var tickChan <-chan tick.Tick

	// initOHLCV(bt.BarsNeeded)
	// activeOrders = make([]order.Order, 0)
	// activePositions = make([]position.Position, 0)

	// // compute first ATR
	// var firstBarTime time.Time
	// if bt.ATRperiod != 0 {
	// 	quotes := make([]quote.Quote, 0, bt.ATRperiod)
	// 	for len(quotes) != bt.ATRperiod {
	// 		quotes = append(quotes, <-quoteChan)
	// 	}
	// 	sum := 0.
	// 	for _, q := range quotes {
	// 		sum += q.High - q.Low
	// 	}
	// 	ATR = sum / float64(len(quotes))
	// 	firstBarTime = quotes[len(quotes)-1].Time.Add(tfDuration)
	// }

	// // create tick channel
	// if bt.HighOrderResolution {
	// 	tickChan = s.tickService.Get(ctx, bt.Symbol, bt.StartDate, bt.EndDate)
	// } else {
	// 	quoteChan, tickChan = ticksFromQuotes(quoteChan, bt.Timeframe)
	// }

	// // pull unnecessary ticks
	// t := <-tickChan
	// for t.Time.Before(firstBarTime) {
	// 	t = <-tickChan
	// }
	// nextBarStart := nextBarStart(t.Time, bt.Timeframe)
	// // fmt.Printf("next bar start: %v\n", nextBarStart)
	// for t := range tickChan {

	// 	if !t.Time.Before(nextBarStart) {
	// 		nextBarStart = nextBarStart.Add(tfDuration)

	// 		q, ok := <-quoteChan
	// 		if !ok {
	// 			log.Info("Quote channel has been closed.")
	// 			break
	// 		}
	// 		updateQuotes(bt.BarsNeeded, q)

	// 		// Update ATR
	// 		if bt.ATRperiod != 0 {
	// 			ATR = ((ATR * float64(bt.ATRperiod-1)) + (High(0) - Low(0))) / float64(bt.ATRperiod)
	// 		}

	// 		bt.Strategy.Implementation.OnQuote(ctx, q)
	// 		// fmt.Printf("Quote: %v\n", Quote(0))
	// 		// fmt.Printf("ATR: %v\n", ATR)
	// 	}
	// 	// fmt.Printf("tick: %v\n", t)

	// 	// exchange logic

	// 	for _, o := range activeOrders {
	// 		if o.Direction == order.Long {
	// 			if o.AbovePrice {
	// 				switch o.Type {
	// 				case order.Market:
	// 				case order.TrailingStop:
	// 					panic("not implemented")
	// 				default: // price has to be reached
	// 					if t.Price <= o.Price {
	// 						// execute
	// 					}
	// 				}
	// 			} else {
	// 				switch o.Type {
	// 				case order.Market:
	// 				case order.TrailingStop:
	// 					panic("not implemented")
	// 				default: // price has to be reached
	// 					if t.Price <= o.Price {
	// 						// execute
	// 					}
	// 				}
	// 			}
	// 		} else {

	// 		}

	// 		// if p.Status == position.Created && p.MainOrder.IsTouched(q) {
	// 		// 	p, err = s.positionService.Open(ctx, p, p.MainOrder.Size)
	// 		// 	if err != nil {
	// 		// 		return err
	// 		// 	}

	// 		// 	bt.Strategy.Implementation.OnOrder(ctx, p.MainOrder)
	// 		// 	bt.Strategy.Implementation.OnPosition(ctx, *p)

	// 		// } else if p.Status == position.Open {
	// 		// 	if p.TakeOrder.IsTouched(q) {
	// 		// 		p, err = s.positionService.TakeProfit(ctx, *p, p.MainOrder.Size)
	// 		// 		if err != nil {
	// 		// 			return err
	// 		// 		}

	// 		// 		bt.Strategy.Implementation.OnOrder(ctx, p.TakeOrder)
	// 		// 		bt.Strategy.Implementation.OnPosition(ctx, *p)

	// 		// 	} else if p.StopOrder.IsTouched(q) {
	// 		// 		p, err = s.positionService.StopLoss(ctx, *p, p.MainOrder.Size)
	// 		// 		if err != nil {
	// 		// 			return err
	// 		// 		}

	// 		// 		bt.Strategy.Implementation.OnOrder(ctx, p.TakeOrder)
	// 		// 		bt.Strategy.Implementation.OnPosition(ctx, *p)
	// 		// 	}
	// 		// }
	// 	}
	// }

	return fmt.Errorf("not implemented")
}

func (s *service) Get(id string) (*Backtester, error) {
	return s.repository.GetBacktest(id)
}

func ticksFromQuotes(quoteChan <-chan quote.Quote, timeframe int) (<-chan quote.Quote, <-chan tick.Tick) {
	newQuoteChan := make(chan quote.Quote, cap(quoteChan))
	tickChan := make(chan tick.Tick, cap(quoteChan)*timeframe)
	go _ticksFromQuotes(quoteChan, newQuoteChan, tickChan)
	return newQuoteChan, tickChan
}

func _ticksFromQuotes(quoteChan <-chan quote.Quote, newQuoteChan chan<- quote.Quote, tickChan chan<- tick.Tick) {
	for q := range quoteChan {
		newQuoteChan <- q
		if q.High-q.Open < q.Open-q.Low {
			tickChan <- tick.Tick{Time: q.Time, Price: q.High}
			tickChan <- tick.Tick{Time: q.Time, Price: q.Low}
		}
		tickChan <- tick.Tick{Time: q.Time, Price: q.Close}

	}
}

func TR(prevQuote, curQuote quote.Quote) float64 {
	max := curQuote.High - curQuote.Low
	max = math.Max(max, math.Abs(curQuote.High-prevQuote.Close))
	max = math.Max(max, math.Abs(curQuote.Low-prevQuote.Close))
	return max
}

func updateQuotes(barsNeeded int, q quote.Quote) {
	if barsNeeded == 1 {
		Quotes[0] = q
	} else {
		Quotes = append(Quotes, q)
		if len(Quotes) == cap(Quotes) {
			copy(Quotes, Quotes[len(Quotes)-barsNeeded:])
			Quotes = Quotes[:barsNeeded]
		}
	}
}

// if t is alligned with timeframe returns bar started at t, else (t + timeframe)
func nextBarStart(t time.Time, timeframe int) time.Time {
	curBarT := t.Truncate(time.Duration(timeframe) * time.Minute)
	if curBarT == t {
		return t
	}
	return curBarT.Add(time.Duration(timeframe) * time.Minute).Truncate(time.Minute)
}
