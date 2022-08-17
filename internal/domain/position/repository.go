package position

type Repository interface {
	Create(p *Position) (*Position, error)
	Update(posID string, fields map[string]interface{}) error
}
