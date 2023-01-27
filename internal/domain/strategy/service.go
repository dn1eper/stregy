package strategy

type Service interface {
	Save(dto CreateStrategyDTO) error
}

type service struct {
	storage Storage
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) Save(dto CreateStrategyDTO) (err error) {
	_, err = s.storage.SaveStrategy(dto.Name, dto.Implementation)
	return err
}
