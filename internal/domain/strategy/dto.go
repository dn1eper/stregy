package strategy

type CreateStrategyDTO struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	Implementation *string
}
