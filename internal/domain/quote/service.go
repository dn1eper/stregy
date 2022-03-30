package quote

import (
	"context"
	"errors"
	"time"
)

type Service interface {
	GetByInterval(ctx context.Context, symbol string, start, end time.Time, offset, pageSize int) ([]Quote, error)
	Load(ctx context.Context, symbol, filePath, delimiter string) error
	Aggregate(ctx context.Context, quotes []Quote, timeframeMSC int64) ([]Quote, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) GetByInterval(ctx context.Context, symbol string, start, end time.Time, offset, pageSize int) ([]Quote, error) {
	return s.repository.GetByInterval(ctx, symbol, start, end, offset, pageSize)
}

func (s *service) Load(ctx context.Context, symbol, filePath, delimiter string) error {
	return s.repository.Load(ctx, symbol, filePath, delimiter)
}

func (s *service) Aggregate(ctx context.Context, quotes []Quote, timeframeMSC int64) ([]Quote, error) {
	quoteTime := quotes[0].Time
	if quoteTime.UnixNano()%1000000 != 0 {
		return nil, errors.New("quotes are not alligned to miliseconds")
	}
	inputTimeframeMSC := quotes[1].Time.Sub(quotes[0].Time).Milliseconds()
	if inputTimeframeMSC > timeframeMSC {
		return nil, errors.New("base timeframe is bigger than required")
	}
	if inputTimeframeMSC == timeframeMSC {
		return quotes, nil
	}

	newQuotes := make([]Quote, 0)
	period := time.Duration(timeframeMSC * 1000000)

	nextQuoteTime := quoteTime.Add(period)
	open := quotes[0].Open
	high := quotes[0].High
	low := quotes[0].Low

	for idx, quote := range quotes {

		if quote.Time == nextQuoteTime || quote.Time.After(nextQuoteTime) {
			nextQuoteTime = nextQuoteTime.Add(period)
			newQuotes = append(newQuotes, Quote{Open: open, High: high, Low: low, Close: quotes[idx-1].Close})
			open = quote.Open
			high = quote.High
			low = quote.Low
		} else {
			if quote.High > high {
				high = quote.High
			}
			if quote.Low < low {
				low = quote.Low
			}
		}
	}

	return newQuotes, nil
}
