package symbol

type Repository interface {
	Create(s Symbol) (*Symbol, error)
	GetByName(name string) (*Symbol, error)
}
