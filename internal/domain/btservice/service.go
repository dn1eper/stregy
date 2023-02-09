package btservice

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path"
	"stregy/internal/domain/bt"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/quote"
	strategy1 "stregy/internal/domain/strategy"
	"stregy/internal/domain/symbol"
	"stregy/internal/domain/tick"
	strategy "stregy/local/strategies/strat1"
	"stregy/pkg/logging"
	"stregy/pkg/utils"
)

type Service interface {
	Create(dto BacktestDTO) (*bt.Backtester, error)
	Launch(bt *bt.Backtester) error
	Run() error
}

type service struct {
	repository        Repository
	tickService       tick.Service
	quoteService      quote.Service
	exgAccService     exgaccount.Service
	symbolService     symbol.Service
	accHistoryService bt.AccountHistoryReport

	stratexecProjectPath string
}

func NewService(
	repository Repository,
	tickService tick.Service,
	quoteService quote.Service,
	exgAccService exgaccount.Service,
	symbolService symbol.Service,
	accHistoryService bt.AccountHistoryReport,
) Service {
	return &service{
		repository:        repository,
		tickService:       tickService,
		quoteService:      quoteService,
		exgAccService:     exgAccService,
		symbolService:     symbolService,
		accHistoryService: accHistoryService,
	}
}

func (s *service) Create(dto BacktestDTO) (*bt.Backtester, error) {
	bt := bt.Backtester{
		StrategyName: dto.StrategyName,
		StartTime:    dto.StartDate,
		EndTime:      dto.EndDate,
		Symbol:       symbol.Symbol{Name: dto.SymbolName},
		TimeframeSec: dto.TimeframeSec,
		Status:       bt.Created,
	}
	return s.repository.Create(bt)
}

func (s *service) Launch(backtester *bt.Backtester) (err error) {
	// check strategy exists
	wd, _ := os.Getwd()
	strategyFilePath := path.Join(wd, "local", "strategies", backtester.StrategyName, "strategy.go")
	if _, err := os.Stat(strategyFilePath); err != nil {
		return errors.New("strategy not found")
	}

	// import strategy needed
	filePath := path.Join(wd, "internal", "domain", "btservice", "service.go")
	importLine := "\tstrategy \"stregy/local/strategies/defaultstrat\""
	newImportLine := fmt.Sprintf("\tstrategy \"stregy/local/strategies/%s\"", backtester.StrategyName)
	err = utils.ReplaceFirstLineInFile(filePath, importLine, newImportLine)
	if err != nil {
		return err
	}

	// run
	go func() {
		executableName := fmt.Sprintf("%s.exe", backtester.ID)
		cmd := exec.Command("go", "build", "-o", executableName, "cmd/main.go")
		err = cmd.Run()
		utils.ReplaceFirstLineInFile(filePath, newImportLine, importLine)
		if err != nil {
			logging.GetLogger().Error(fmt.Sprintf("backtest build error: %s", err.Error()))
			return
		}

		executablePath := fmt.Sprintf("%s\\%s", wd, executableName)
		cmd = exec.Command(executablePath, "--backtest", backtester.ID)
		defer func() {
			os.Remove(executablePath)
		}()
		err = cmd.Run()
		if err != nil {
			logging.GetLogger().Error(fmt.Sprintf("backtest run error: %s", err.Error()))
		}
	}()

	return nil
}

func (s *service) Run() (err error) {
	serviceLogger := logging.GetLogger()
	defer func() {
		if err != nil {
			serviceLogger.Error(err.Error())
		}
	}()

	backtestID, reportLocation, err := parseArgs()
	if err != nil {
		return err
	}

	backtester, err := s.repository.GetBacktest(backtestID)
	if err != nil {
		return err
	}
	backtester.AccountHistoryService = s.accHistoryService
	backtester.Symbol = *s.getSymbol(backtester.Symbol.Name)

	var strat strategy1.Strategy = strategy.NewStrategy(backtester)

	// backtest
	serviceLogger.Info(fmt.Sprintf("running backtest with strategy %v on period [%s; %s]", strat.Name(), backtester.StartTime.Format("2006-01-02 15:04:05"), backtester.EndTime.Format("2006-01-02 15:04:05")))
	quotes, firstQuote := s.quoteService.Get(backtester.Symbol.Name, backtester.StartTime, backtester.EndTime, backtester.TimeframeSec)
	backtester.BacktestOnQuotes(strat, quotes, firstQuote)
	backtester.CreateReport(reportLocation)

	return err
}

func (s *service) getSymbol(name string) *symbol.Symbol {
	smbl, _ := s.symbolService.GetByName(name)
	if s == nil {
		smbl = &symbol.Symbol{Name: name, Precision: 6}
	}

	return smbl
}

func parseArgs() (backtestID string, reportLocation string, err error) {
	if len(os.Args) < 2 {
		return "", "", errors.New("backtest id not provided")
	}
	backtestID = os.Args[2]

	if len(os.Args) >= 3 {
		reportLocation = os.Args[3]
	}

	return backtestID, reportLocation, nil
}
