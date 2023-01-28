package backtester

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	strategy1 "stregy/internal/domain/strategy"
	"stregy/internal/domain/tick"
	strategy "stregy/local/strategies/defaultstrat"
	"stregy/pkg/logging"
	"stregy/pkg/utils"
)

type Service interface {
	Create(dto BacktestDTO) (*Backtest, error)
	Launch(bt *Backtest) error
	Run() error
}

type service struct {
	repository      Repository
	tickService     tick.Service
	quoteService    quote.Service
	exgAccService   exgaccount.Service
	positionService position.Service

	stratexecProjectPath string

	positions []*position.Position
}

func NewService(
	repository Repository,
	tickService tick.Service,
	quoteService quote.Service,
	exgAccService exgaccount.Service,
	positionService position.Service,
) Service {
	return &service{
		repository:      repository,
		tickService:     tickService,
		quoteService:    quoteService,
		exgAccService:   exgAccService,
		positionService: positionService,
	}
}

func (s *service) Create(dto BacktestDTO) (*Backtest, error) {
	bt := Backtest{
		StrategyName:        dto.StrategyName,
		StartTime:           dto.StartDate,
		EndTime:             dto.EndDate,
		Symbol:              dto.Symbol,
		Timeframe:           dto.Timeframe,
		HighOrderResolution: dto.HighOrderResolution,
		Status:              Created,
	}
	return s.repository.Create(bt)
}

func (s *service) Launch(bt *Backtest) (err error) {
	// check strategy exists
	wd, _ := os.Getwd()
	strategyFilePath := path.Join(wd, "local", "strategies", bt.StrategyName, "strategy.go")
	if _, err := os.Stat(strategyFilePath); err != nil {
		return errors.New("strategy not found")
	}

	// import strategy needed
	filePath := path.Join(wd, "internal", "domain", "backtester", "service.go")
	importLine := "\tstrategy \"stregy/local/strategies/defaultstrat\""
	newImportLine := fmt.Sprintf("\tstrategy \"stregy/local/strategies/%s\"", bt.StrategyName)
	err = utils.ReplaceFirstLineInFile(filePath, importLine, newImportLine)
	if err != nil {
		return err
	}
	// defer utils.ReplaceFirstLineInFile(filePath, newImportLine, importLine)

	// run
	go func() {
		cmd := exec.Command("go", "run", "cmd/main.go", "--backtest", bt.Id)
		err = cmd.Run()
		if err != nil {
			logging.GetLogger().Error(fmt.Sprintf("backtest run error: %s", err.Error()))
		}
	}()

	return nil
}

// Update Quotes in sync, mb prep new Quotes in seperate goroutine
func (s *service) Run() (err error) {
	logger := logging.GetLogger()

	if len(os.Args) < 2 {
		return errors.New("backtest id not provided")
	}

	_, err = s.repository.GetBacktest(os.Args[2])
	if err != nil {
		return err
	}

	// backtest
	var strat strategy1.Strategy = &strategy.Strategy{}
	logger.Debug(fmt.Sprintf("running backtest with strategy %v", strat))

	return err
}
