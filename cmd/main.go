package main

import (
	"stregy/internal/app"
	"stregy/internal/config"
	"stregy/pkg/logging"

	log "github.com/sirupsen/logrus"
)

func main() {
	// entry point
	log.Info("config initializing")
	cfg := config.GetConfig()
	logging.Init(cfg.LogLevel, "logs/all.log")

	app.Run(cfg)
}
