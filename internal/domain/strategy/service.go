package strategy

import (
	"context"
	"encoding/base64"
	"os"
	"path/filepath"
	"stregy/internal/domain/user"
	"stregy/pkg/utils"
)

type Service interface {
	GetByUUID(ctx context.Context, id string) (*Strategy, error)
	Create(ctx context.Context, strategy CreateStrategyDTO, user *user.User) (*Strategy, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(ctx context.Context, dto CreateStrategyDTO, user *user.User) (strategy *Strategy, err error) {
	strategy = &Strategy{Name: dto.Name, Description: dto.Description}

	strategy, err = s.repository.Create(ctx, *strategy)
	if err != nil {
		return nil, err
	}

	dec, _ := base64.StdEncoding.DecodeString(dto.Implementation)
	dirpath, _ := utils.CreateStratRepo(user.ID, strategy.ID)
	f, err := os.Create(filepath.Join(dirpath, "strategy"))
	defer f.Close()
	if err != nil {
		return nil, err
	}
	f.Write(dec)

	return strategy, nil
}

func (s *service) GetByUUID(ctx context.Context, uuid string) (strategy *Strategy, err error) {
	strategy, err = s.repository.GetOne(ctx, uuid)
	if err != nil {
		return nil, err
	}

	//TODO: get strategy implementation

	return strategy, nil
}
