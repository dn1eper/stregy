package main

import (
	"stregy/internal/app"
	"stregy/internal/config"
	"stregy/pkg/logging"
)

func main() {
	// entry point
	logging.Init()
	logger := logging.GetLogger()

	logger.Info("config initializing")
	cfg := config.GetConfig()

	app.Run(cfg)
}
