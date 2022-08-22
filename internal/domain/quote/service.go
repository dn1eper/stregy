package quote

import (
	"context"
	"fmt"
	"time"
)

type Service interface {
	Get(ctx context.Context, symbol string, start, end time.Time, timeframe int) <-chan Quote
	Load(symbol, filePath, delimiter string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Get(ctx context.Context, symbol string, start, end time.Time, timeframe int) <-chan Quote {
	ch := make(chan Quote, 256)
	go quoteGenerator(ctx, ch, s, symbol, start, end, timeframe)
	return ch
}

func quoteGenerator(ctx context.Context, ch chan<- Quote, s *service, symbol string, start, end time.Time, timeframe int) {
	timeframeSec := timeframe * 60
	batchStart := start
	batchEnd := batchStart.AddDate(0, 0, 1)
	if batchEnd.After(end) {
		batchEnd = end
	}
	if 86400%timeframeSec != 0 {
		panic("one day is not a multiple of requested timeframe")
	}

	for true {
		quotes, err := s.repository.GetByInterval(ctx, symbol, batchStart, batchEnd)
		if len(quotes) == 0 {
			break
		}

		quotesAgg, err := Aggregate(ctx, quotes, timeframeSec)
		if err != nil {
			panic(fmt.Sprintf("error aggregating quotes: %v\n", err))
		}

		for _, quote := range quotesAgg {
			ch <- quote
		}

		batchStart = batchEnd
		batchEnd = batchStart.AddDate(0, 0, 1)
		if batchEnd.After(end) {
			batchEnd = end
		}
	}
	close(ch)
}

func (s *service) Load(symbol, filePath, delimiter string) error {
	return s.repository.Load(symbol, filePath, delimiter)
}
