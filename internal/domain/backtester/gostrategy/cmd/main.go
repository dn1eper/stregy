package main

import (
	"stregy/internal/config"
	"stregy/internal/domain/backtester/gostrategy/app"
	"stregy/pkg/logging"

	log "github.com/sirupsen/logrus"
)

func main() {
	// entry point
	log.Info("config initializing")
	cfg := config.GetConfig()
	logging.Init(cfg.LogLevel)

	app.Run(cfg)
}
