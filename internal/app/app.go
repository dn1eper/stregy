package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"stregy/internal/composites"
	"stregy/internal/config"
	"stregy/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Run(cfg *config.Config) {
	logger := logging.GetLogger()

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("pgorm composite initializing")
	PGormComposite, err := composites.NewPGormComposite(context.Background(), cfg.PosgreSQL.Host, cfg.PosgreSQL.Port, cfg.PosgreSQL.Username, cfg.PosgreSQL.Password, cfg.PosgreSQL.Database)
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}

	logger.Info("user composite initializing")
	userComposite, err := composites.NewUserComposite(PGormComposite)
	if err != nil {
		logger.Fatal("author composite failed")
	}
	userComposite.Handler.Register(router)

	logger.Info("listener initializing")
	listener, err := net.Listen("tcp", "")
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
