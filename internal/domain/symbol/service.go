package symbol

type Service interface {
	Exists(name string) bool
	Create(name string) (*Symbol, error)
	GetAll() ([]Symbol, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) Exists(name string) bool {
	return s.repository.Exists(name)
}
func (s *service) Create(name string) (*Symbol, error) {
	return s.repository.Create(name)
}

func (s *service) GetAll() ([]Symbol, error) {
	return s.repository.GetAll()
}
