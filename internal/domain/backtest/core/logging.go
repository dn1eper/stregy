package core

func (b *Backtest) Print(s string) {
	b.logger.Print(s)
}

func (b *Backtest) Printf(format string, v ...interface{}) {
	b.logger.Printf(format, v...)
}
