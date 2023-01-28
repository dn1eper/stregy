package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"stregy/internal/composites"
	"stregy/internal/config"
	"stregy/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

type AppMode int

const (
	Server AppMode = iota
	Backtest
)

var appMode AppMode

func SetAppMode() {
	if len(os.Args) < 2 {
		appMode = Server
		return
	}

	switch os.Args[1] {
	case "--backtest":
		appMode = Backtest
	default:
		panic(fmt.Errorf("unknown app mode: %s", os.Args[1]))
	}
}

func Run(cfg *config.Config) {
	logger := logging.GetLogger()

	SetAppMode()

	logger.Info("pgorm composite initialization")
	pgormComposite, err := composites.NewPGormComposite(cfg.PosgreSQL.Host, cfg.PosgreSQL.Port, cfg.PosgreSQL.Username, cfg.PosgreSQL.Password, cfg.PosgreSQL.Database)
	if err != nil {
		logger.Fatal("pgorm composite failed")
	}

	logger.Info("user composite initialization")
	userComposite, err := composites.NewUserComposite(pgormComposite)
	if err != nil {
		logger.Fatal("user composite failed")
	}

	logger.Info("quote composite initialization")
	quoteComposite, err := composites.NewQuoteComposite(pgormComposite)
	if err != nil {
		logger.Fatal("quote composite failed")
	}

	logger.Info("tick composite initialization")
	tickComposite, err := composites.NewTickComposite(pgormComposite)
	if err != nil {
		logger.Fatal("tick composite failed")
	}

	logger.Info("strategy composite initialization")
	strategyComposite, err := composites.NewStrategyComposite(pgormComposite, userComposite.Service)
	if err != nil {
		logger.Fatal("strategy composite failed")
	}

	logger.Info("exchange account composite initialization")
	exgAccountComposite, err := composites.NewExchangeAccountComposite(pgormComposite, userComposite.Service)
	if err != nil {
		logger.Fatal("exchange account composite failed")
	}

	logger.Info("symbol composite initialization")
	_, err = composites.NewSymbolComposite(pgormComposite)
	if err != nil {
		logger.Fatal("symbol composite failed")
	}

	logger.Info("order composite initialization")
	orderComposite, err := composites.NewOrderComposite()
	if err != nil {
		logger.Fatal("order composite failed")
	}

	logger.Info("position composite initialization")
	positionComposite, err := composites.NewPositionComposite(pgormComposite, orderComposite.Service)
	if err != nil {
		logger.Fatal("position composite failed")
	}

	logger.Info("backtester composite initialization")
	backtesterComposite, err := composites.NewBacktesterComposite(
		pgormComposite, exgAccountComposite.Service,
		strategyComposite.Service, userComposite.Service,
		tickComposite.Service, quoteComposite.Service,
		positionComposite.Service,
	)
	if err != nil {
		logger.Fatal("backtester composite failed")
	}

	switch appMode {
	case Server:
		StartServer(userComposite, strategyComposite, exgAccountComposite, backtesterComposite)
	case Backtest:
		err := backtesterComposite.Service.Run()
		if err != nil {
			logger.Error(err.Error())
		}
	}
}

func StartServer(
	userComposite *composites.UserComposite,
	strategyComposite *composites.StrategyComposite,
	exgAccountComposite *composites.ExchangeAccountComposite,
	backtesterComposite *composites.BacktesterComposite) {
	logger := logging.GetLogger()

	logger.Info("router initialization")
	router := httprouter.New()

	userComposite.Handler.Register(router)
	strategyComposite.Handler.Register(router)
	exgAccountComposite.Handler.Register(router)
	backtesterComposite.Handler.Register(router)

	logger.Info("listener initialization")
	cfg := config.GetConfig()
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v", cfg.Listen.BindIP, cfg.Listen.Port))
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatalln(server.Serve(listener))
}
