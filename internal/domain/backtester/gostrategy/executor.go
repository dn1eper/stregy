package gostrategy

import (
	"context"
	"fmt"
	"os/exec"
	"stregy/internal/domain/backtester"
	"stregy/pkg/utils"

	log "github.com/sirupsen/logrus"
)

const stregyCopyPath string = "stregycopy"

type executor struct {
}

func NewExecutor() backtester.Executor {
	return &executor{}
}

func (e *executor) Execute(ctx context.Context, b *backtester.Backtester) error {
	strat_rout := "repository/strategies/" + b.Strategy.ID + "/strategy.zip"

	// unzip strategy
	err := utils.Unzip(strat_rout, stregyCopyPath+"/internal/domain/backtester/gostrategy/strategy")
	if err != nil {
		log.Error(err)
		return err
	}

	// run
	log.Info("starting backtest")
	cmd := exec.Command("go", "run", stregyCopyPath+"/internal/domain/backtester/gostrategy/cmd/main.go", b.ID)
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return err
	}
	go func() {
		if err := cmd.Wait(); err != nil {
			log.Error(fmt.Sprintf("%v: %v", err, cmd.Stdout))
		}
	}()

	return nil
}
