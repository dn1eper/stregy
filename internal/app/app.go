package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"stregy/internal/composites"
	"stregy/internal/config"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/julienschmidt/httprouter"
)

func Run(cfg *config.Config) {
	log.Info("router initialization")
	router := httprouter.New()

	log.Info("pgorm composite initialization")
	pgormComposite, err := composites.NewPGormComposite(context.Background(), cfg.PosgreSQL.Host, cfg.PosgreSQL.Port, cfg.PosgreSQL.Username, cfg.PosgreSQL.Password, cfg.PosgreSQL.Database)
	if err != nil {
		log.Fatal("pgorm composite failed")
	}

	log.Info("user composite initialization")
	userComposite, err := composites.NewUserComposite(pgormComposite)
	if err != nil {
		log.Fatal("user composite failed")
	}
	userComposite.Handler.Register(router)

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
	// tickComposite.Service.Load(context.TODO(), )

	log.Info("strategy composite initialization")
	strategyComposite, err := composites.NewStrategyComposite(pgormComposite, userComposite.Service)
	if err != nil {
		log.Fatal("strategy composite failed")
	}
	strategyComposite.Handler.Register(router)

	log.Info("exchange account composite initialization")
	exgAccountComposite, err := composites.NewExchangeAccountComposite(pgormComposite, userComposite.Service)
	if err != nil {
		log.Fatal("exchange account composite failed")
	}
	exgAccountComposite.Handler.Register(router)

	log.Info("symbol composite initialization")
	_, err = composites.NewSymbolComposite(pgormComposite)
	if err != nil {
		log.Fatal("symbol composite failed")
	}

	log.Info("order composite initialization")
	orderComposite, err := composites.NewOrderComposite()
	if err != nil {
		log.Fatal("order composite failed")
	}

	log.Info("position composite initialization")
	positionComposite, err := composites.NewPositionComposite(pgormComposite, orderComposite.Service)
	if err != nil {
		log.Fatal("position composite failed")
	}

	log.Info("backtester composite initialization")
	backtesterComposite, err := composites.NewBacktesterComposite(
		pgormComposite, exgAccountComposite.Service,
		strategyComposite.Service, userComposite.Service,
		tickComposite.Service, quoteComposite.Service,
		positionComposite.Service,
	)
	if err != nil {
		log.Fatal("backtester composite failed")
	}
	backtesterComposite.Handler.Register(router)

	log.Info("listener initialization")
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
