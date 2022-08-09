package tick

import (
	"context"
	"time"
)

type Service interface {
	Get(ctx context.Context, symbol string, start, end time.Time) chan Tick
	Load(symbol, filePath, delimiter string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s service) Load(symbol, filePath, delimiter string) error {
	return s.repository.Load(symbol, filePath, delimiter)
}

func (s *service) Get(ctx context.Context, symbol string, start, end time.Time) chan Tick {
	ch := make(chan Tick, 10000)
	go tickGenerator(ctx, ch, s, symbol, start, end)

	return ch
}

func tickGenerator(ctx context.Context, ch chan<- Tick, s *service, symbol string, start, end time.Time) {
	batchStart := start
	batchEnd := batchStart.AddDate(0, 0, 1)
	if batchEnd.After(end) {
		batchEnd = end
	}

	for true {
		ticks, err := s.repository.GetByInterval(ctx, symbol, batchStart, batchEnd)
		if err != nil {
			panic(err)
		}
		if len(ticks) == 0 {
			break
		}

		for _, tick := range ticks {
			ch <- tick
		}

		batchStart = batchEnd
		batchEnd = batchStart.AddDate(0, 0, 1)
		if batchEnd.After(end) {
			batchEnd = end
		}
	}
	close(ch)
}