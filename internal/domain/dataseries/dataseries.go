package dataseries

import (
	"stregy/internal/domain/quote"
)

// DataSeries should be syncronized before any user code is run
type DataSeries struct {
	Quotes     []quote.Quote
	ATR        []float64
	barsNeeded int
}

func NewDataSeries(barsNeeded int) DataSeries {
	var quotes []quote.Quote
	if barsNeeded == 1 {
		quotes = make([]quote.Quote, 1)
	} else {
		quotes = make([]quote.Quote, 0, barsNeeded*2)
	}

	return DataSeries{
		Quotes:     quotes,
		barsNeeded: barsNeeded,
	}
}

func (ds *DataSeries) Add(quote *quote.Quote) {
	if ds.barsNeeded == 1 {
		ds.Quotes[0] = *quote
	} else {
		ds.Quotes = append(ds.Quotes, *quote)
		if len(ds.Quotes) == cap(ds.Quotes) {
			copy(ds.Quotes, ds.Quotes[len(ds.Quotes)-ds.barsNeeded:])
			ds.Quotes = ds.Quotes[:ds.barsNeeded]
		}
	}
}
