package app

import (
	"context"
	"stregy/internal/composites"
	"stregy/internal/config"
	"stregy/pkg/logging"

	"github.com/julienschmidt/httprouter"
)

func Run(cfg *config.Config) {
	logger := logging.GetLogger()

	logger.Info("creater router")
	router := httprouter.New()

	logger.Info("postgresql composite initializing")
	postgreComposite, err := composites.NewPostgreSQLComposite(context.Background(), cfg.PosgreSQL.Host, cfg.PosgreSQL.Port, cfg.PosgreSQL.Username, cfg.PosgreSQL.Password, cfg.PosgreSQL.Database)
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}

	logger.Info("user composite initializing")
	userComposite, err := composites.NewUserComposite(postgreComposite)
	if err != nil {
		logger.Fatal("author composite failed")
	}
	userComposite.Handler.Register(router)

}
