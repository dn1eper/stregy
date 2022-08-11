package quote

import (
	"context"
	"errors"
	"time"
)

func Aggregate(ctx context.Context, quotes []Quote, timeframeSec int) ([]Quote, error) {
	quoteTime := quotes[0].Time
	inputTimeframeMsc := quotes[1].Time.Sub(quotes[0].Time).Milliseconds()
	if inputTimeframeMsc > int64(timeframeSec)*1000 {
		return nil, errors.New("base timeframe is bigger than one you asked to aggregate to")
	}
	if inputTimeframeMsc == int64(timeframeSec)*1000 {
		return quotes, nil
	}

	newQuotes := make([]Quote, 0)
	period := time.Duration(timeframeSec * 1000000000)

	nextQuoteTime := quoteTime.Add(period)
	open := quotes[0].Open
	high := quotes[0].High
	low := quotes[0].Low

	for idx, quote := range quotes {
		if quote.Time.Before(nextQuoteTime) && idx != len(quotes)-1 {
			if high < quote.High {
				high = quote.High
			}
			if quote.Low < low {
				low = quote.Low
			}
		} else {
			newQuotes = append(newQuotes, Quote{Time: quoteTime, Open: open, High: high, Low: low, Close: quotes[idx-1].Close})
			quoteTime = nextQuoteTime
			nextQuoteTime = nextQuoteTime.Add(period)
			open = quote.Open
			high = quote.High
			low = quote.Low
		}
	}

	return newQuotes, nil
}
