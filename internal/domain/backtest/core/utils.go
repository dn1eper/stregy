package core

import (
	"os"
	"path"
)

func (b *Backtest) getDefaultReportPath() string {
	wd, _ := os.Getwd()
	reportDir := path.Join(wd, "reports")
	os.Mkdir(reportDir, os.ModePerm)
	return path.Join(reportDir, b.ID+".csv")
}
