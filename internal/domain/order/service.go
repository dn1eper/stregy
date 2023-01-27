package order

type Service interface {
	Create() (Order, error)
	Open(id string) error
	Fill(id string, size float64) (Order, error)
	Submit(id string) (Order, error)
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Create() (Order, error) {
	panic("not implemented")
}

func (s *service) Open(id string) error {
	panic("not implemented")
}

func (s *service) Fill(id string, size float64) (Order, error) {
	panic("not implemented")
}

func (s *service) Submit(id string) (Order, error) {
	panic("not implemented")
}
