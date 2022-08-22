package main

import (
	"fmt"
	"os"
	"stregy/internal/config"
	"stregy/internal/domain/backtester/gostrategy/app"
	"stregy/pkg/logging"

	log "github.com/sirupsen/logrus"
)

func main() {
	// entry point
	log.Info("config initializing")
	cfg := config.GetConfig()
	stratexecID := os.Args[1]
	logDirPath := os.Args[2]
	logging.Init(cfg.LogLevel, fmt.Sprintf("%v/%v.log", logDirPath, stratexecID))

	app.Run(cfg)
}
