package symbol

type Service interface {
	Create(symbol Symbol) (*Symbol, error)
	GetByName(name string) (*Symbol, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Create(symbol Symbol) (*Symbol, error) {
	return s.repository.Create(symbol)
}

func (s *service) GetByName(name string) (*Symbol, error) {
	return s.repository.GetByName(name)
}
