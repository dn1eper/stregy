package bt

func (b *Backtester) Print(s string) {
	b.logger.Print(s)
}

func (b *Backtester) Printf(format string, v ...interface{}) {
	b.logger.Printf(format, v...)
}
