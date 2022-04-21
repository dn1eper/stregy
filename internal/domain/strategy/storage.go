package strategy

type Storage interface {
	SaveStrategy(implementation *string, strategyID string) error
}
