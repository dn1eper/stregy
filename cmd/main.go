package main

import (
	"stregy/internal/app"
	"stregy/internal/config"
	"stregy/pkg/logging"

	"github.com/sirupsen/logrus"
)

func main() {
	logging.Init(logrus.DebugLevel)

	cfg := config.GetConfig()

	app.Run(cfg)
}
