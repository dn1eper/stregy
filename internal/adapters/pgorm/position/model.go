package position

import (
	"stregy/internal/adapters/pgorm/order"

	"github.com/google/uuid"
)

type Position struct {
	ID          uuid.UUID   `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Num         int64       `gorm:"type:integer"`
	MainOrder   order.Order `gorm:"foreignKey:MainOrderID"`
	MainOrderID uuid.UUID
	Orders      []order.Order
}
