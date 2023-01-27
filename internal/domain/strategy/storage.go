package strategy

type Storage interface {
	SaveStrategy(name string, implementation *string) (string, error)
}
