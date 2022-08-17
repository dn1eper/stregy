package order

type Repository interface {
	Create(o *Order, stratexecID string) (*Order, error)
	Get(orderID string) *Order
	Update(orderID string, fields map[string]interface{}) error
}
