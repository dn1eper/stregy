package app

import (
	"context"
	"os"
	"stregy/internal/composites"
	"stregy/internal/config"
	"stregy/internal/domain/backtester/gostrategy"

	log "github.com/sirupsen/logrus"
)

func Run(cfg *config.Config) {
	log.Info("pgorm composite initialization")
	pgormComposite, err := composites.NewPGormComposite(context.Background(), cfg.PosgreSQL.Host, cfg.PosgreSQL.Port, cfg.PosgreSQL.Username, cfg.PosgreSQL.Password, cfg.PosgreSQL.Database)
	if err != nil {
		log.Fatal("pgorm composite failed: ", err)
	}

	log.Info("user composite initialization")
	userComposite, err := composites.NewUserComposite(pgormComposite)
	if err != nil {
		log.Fatal("user composite failed")
	}

	log.Info("quote composite initialization")
	quoteComposite, err := composites.NewQuoteComposite(pgormComposite)
	if err != nil {
		log.Fatal("quote composite failed")
	}

	log.Info("tick composite initialization")
	tickComposite, err := composites.NewTickComposite(pgormComposite)
	if err != nil {
		log.Fatal("tick composite failed")
	}

	log.Info("strategy composite initialization")
	strategyComposite, err := composites.NewStrategyComposite(pgormComposite, userComposite.Service)
	if err != nil {
		log.Fatal("strategy composite failed")
	}

	log.Info("exchange account composite initialization")
	exgAccountComposite, err := composites.NewExchangeAccountComposite(pgormComposite, userComposite.Service)
	if err != nil {
		log.Fatal("exchange account composite failed")
	}

	log.Info("symbol composite initialization")
	_, err = composites.NewSymbolComposite(pgormComposite)
	if err != nil {
		log.Fatal("symbol composite failed")
	}

	log.Info("order composite initialization")
	orderComposite, err := composites.NewOrderComposite(pgormComposite)
	if err != nil {
		log.Fatal("order composite failed")
	}

	log.Info("position composite initialization")
	positionComposite, err := composites.NewPositionComposite(pgormComposite, orderComposite.Service)
	if err != nil {
		log.Fatal("position composite failed")
	}

	backtestExecutor := gostrategy.NewExecutor()
	log.Info("backtester composite initialization")
	backtesterComposite, err := composites.NewBacktesterComposite(
		pgormComposite, exgAccountComposite.Service,
		strategyComposite.Service, userComposite.Service,
		tickComposite.Service, quoteComposite.Service,
		positionComposite.Service,
		backtestExecutor,
	)
	if err != nil {
		log.Fatal("backtester composite failed")
	}

	// execute
	if len(os.Args) < 2 {
		log.Error("no backtester id provided")
		return
	}
	backtesterID := os.Args[1]
	bt, err := backtesterComposite.Service.Get(backtesterID)
	err = backtesterComposite.Service.Run(context.Background(), bt)
	if err != nil {
		log.Error(err)
		return
	}
}
