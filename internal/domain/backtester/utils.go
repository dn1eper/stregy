package backtester

import (
	"stregy/internal/domain/quote"
	"stregy/internal/domain/tick"
)

func ticksFromQuotes(quoteChan <-chan quote.Quote, timeframe int) (<-chan quote.Quote, <-chan tick.Tick) {
	newQuoteChan := make(chan quote.Quote, cap(quoteChan))
	tickChan := make(chan tick.Tick, cap(quoteChan)*timeframe)
	go _ticksFromQuotes(quoteChan, newQuoteChan, tickChan)
	return newQuoteChan, tickChan
}

func _ticksFromQuotes(quoteChan <-chan quote.Quote, newQuoteChan chan<- quote.Quote, tickChan chan<- tick.Tick) {
	for q := range quoteChan {
		newQuoteChan <- q
		if q.High-q.Open < q.Open-q.Low {
			tickChan <- tick.Tick{Time: q.Time, Price: q.High}
			tickChan <- tick.Tick{Time: q.Time, Price: q.Low}
		}
		tickChan <- tick.Tick{Time: q.Time, Price: q.Close}

	}
}
