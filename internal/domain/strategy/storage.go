package strategy

type Storage interface {
	SaveStrategy(implementation, userID, strategyID string) error
}
