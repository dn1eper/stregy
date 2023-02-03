package position

import (
	"stregy/internal/adapters/pgorm/order"

	"github.com/google/uuid"
)

type Position struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Num         int64     `gorm:"type:integer"`
	MainOrder   order.Order
	StopOrder   order.Order
	TakeOrder   order.Order
	MainOrderId uuid.UUID
	StopOrderId uuid.UUID
	TakeOrderId uuid.UUID
}
