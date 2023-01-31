package bt

import (
	"fmt"
	"os"
	"path"
)

type Logger interface {
	Print(s string)
	Printf(format string, v ...interface{})
}

type logger struct {
	file *os.File
}

var loggerInstance logger

func InitLogger(name string) {
	wd, _ := os.Getwd()
	dir := path.Join(wd, "logs", "backtest")
	os.Mkdir(dir, os.ModePerm)
	fpath := path.Join(dir, name)
	f, err := os.OpenFile(fpath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(fmt.Errorf("could not create log file: %s", err.Error()))
	}

	loggerInstance = logger{file: f}
}

func Print(s string) {
	s = timePrefix() + s
	fmt.Println(s)
	loggerInstance.file.WriteString(s + "\n")
}

func Printf(format string, v ...interface{}) {
	Print(fmt.Sprintf(format, v...))
}

func timePrefix() string {
	return Time.Format("2006-01-02 15:04:05") + ": "
}
