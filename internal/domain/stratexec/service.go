package stratexec

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"stregy/internal/domain/exgaccount"
	"stregy/internal/domain/user"
	"time"
)

type Service interface {
	Create(ctx context.Context, se CreateStrategyExecutionDTO, user *user.User) (*StrategyExecution, error)
}

type service struct {
	repository    Repository
	exgAccService exgaccount.Service
}

func NewService(repository Repository, exgAccService exgaccount.Service) Service {
	return &service{repository: repository, exgAccService: exgAccService}
}

func (s *service) Create(ctx context.Context, dto CreateStrategyExecutionDTO, user *user.User) (strategy *StrategyExecution, err error) {
	userID := s.exgAccService.GetUserID(ctx, dto.ExchangeAccountID)
	if userID != user.ID {
		fmt.Printf("exg account user id: %v, request api key user id: %v\n", userID, user.ID)
		return nil, errors.New("incorrect exchange account id")
	}

	timeframe, _ := strconv.Atoi(dto.Timeframe)
	startTime, _ := time.Parse("2006-01-02", dto.StartTime)
	endTime, _ := time.Parse("2006-01-02", dto.EndTime)
	se := &StrategyExecution{
		StrategyID:        dto.StrategyID,
		ExchangeAccountID: dto.ExchangeAccountID,
		Timeframe:         timeframe,
		Symbol:            dto.Symbol,
		StartTime:         startTime,
		EndTime:           endTime,
	}
	se, err = s.repository.Create(ctx, *se)
	if err != nil {
		return nil, err
	}
	return se, nil
}
