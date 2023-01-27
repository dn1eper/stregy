package strategy

type CreateStrategyDTO struct {
	Name           string `json:"name"`
	Implementation *string
}
