package symbol

type Repository interface {
	Exists(name string) bool
	Create(name string) (*Symbol, error)
	GetAll() ([]Symbol, error)
}
