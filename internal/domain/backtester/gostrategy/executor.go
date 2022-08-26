package gostrategy

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"stregy/internal/domain/backtester"
	"stregy/pkg/utils"

	log "github.com/sirupsen/logrus"
)

const stregyCopyPath string = "stregycopy/stregy"

type executor struct {
	stregyCopyFullPath string
}

func NewExecutor() (backtester.Executor, error) {
	utils.CreateDir(stregyCopyPath)
	wd, _ := os.Getwd()
	stregyCopyFullPath := filepath.Join(wd, stregyCopyPath)
	err := utils.CopyFolderWin(wd, stregyCopyFullPath,
		[]string{
			".git", ".vscode", "cmd", "logs", "repository",
			"stregycopy", "test", ".gitignore", "main.exe",
			"readme.md", "todo.md",
		})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &executor{stregyCopyFullPath}, nil
}

func (e *executor) Execute(ctx context.Context, b *backtester.Backtester) error {
	strat_rout := "repository/strategies/" + b.Strategy.ID + "/strategy.zip"

	// unzip strategy
	os.RemoveAll(stregyCopyPath + "/user/strategy")
	err := utils.Unzip(strat_rout, stregyCopyPath+"/user/strategy")
	if err != nil {
		log.Error(err)
		return err
	}

	// run
	logDirPath, _ := os.Getwd()
	logDirPath, err = utils.CreateDir(logDirPath, "logs", "stratexec")

	if err != nil {
		log.Error(err)
		return err
	}
	log.Info("starting backtest")
	cmd := exec.Command("go", "run", "internal/domain/backtester/gostrategy/cmd/main.go", b.ID, logDirPath)
	cmd.Dir = e.stregyCopyFullPath
	go func() {
		output, err := cmd.CombinedOutput()
		if err != nil {
			err = fmt.Errorf("%v: %v", cmd, err)
			log.Error(err)
			if len(output) != 0 {
				f, err := os.OpenFile(logDirPath+"/all.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
				f.Write([]byte("----------------------------------------------------\n"))
				if err != nil {
					log.Error(err)
				}
				defer f.Close()
				f.Write(output)
			}
		}
	}()

	return nil
}
