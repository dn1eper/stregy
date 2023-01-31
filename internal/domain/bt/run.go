package bt

import (
	"os"
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
)

var Orders []order.Order
var Positions []position.Position

func RunOnQuotes(backtest *Backtest, strat strategy.Strategy, quotes <-chan quote.Quote, firstQuote quote.Quote) error {
	Time = firstQuote.Time
	InitLogger(backtest.Id + ".log")
	Printf("period [%s; %s]", backtest.StartTime.Format("2006-01-02 15:04:05"), backtest.EndTime.Format("2006-01-02 15:04:05"))

	quoteGen, err := NewQuoteGenerator(strat, backtest.TimeframeSec, firstQuote)
	if err != nil {
		return err
	}

	qCount := 0
	for q := range quotes {
		Time = q.Time

		// debug breaker
		qCount += 1
		if qCount == 3601 {
			break
		}

		quoteGen.OnQuote(q)
	}

	return nil
}

func Terminate() {
	Print("Terminated")
	os.Exit(0)
}
