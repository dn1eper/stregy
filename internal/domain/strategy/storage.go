package strategy

type Storage interface {
	SaveStrategy(implementation *string, userID, strategyID string) error
}
