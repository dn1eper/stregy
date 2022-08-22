package tick

import (
	"context"
	"time"
)

type Service interface {
	GetHistorical(ctx context.Context, symbol string, start, end time.Time) <-chan Tick
	GetLive(symbol string) <-chan Tick
	Upload(symbol, filePath, delimiter string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s service) Upload(symbol, filePath, delimiter string) error {
	return s.repository.Load(symbol, filePath, delimiter)
}

func (s service) GetHistorical(ctx context.Context, symbol string, start, end time.Time) <-chan Tick {
	ch := make(chan Tick, 10000)
	go tickGenerator(ctx, ch, s, symbol, start, end)

	return ch
}

func (s service) GetLive(symbol string) <-chan Tick {
	panic("not implemented")
}
