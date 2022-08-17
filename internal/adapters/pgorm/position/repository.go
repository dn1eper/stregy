package position

import (
	"stregy/internal/domain/position"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) position.Repository {
	return &repository{db: client}
}

func (r *repository) Create(p *position.Position) (*position.Position, error) {
	mainOrderID, _ := uuid.Parse(p.MainOrder.OrderID)
	stopOrderID, _ := uuid.Parse(p.StopOrder.OrderID)
	takeOrderID, _ := uuid.Parse(p.TakeOrder.OrderID)

	positionDB := &Position{
		MainOrderID: mainOrderID,
		StopOrderID: stopOrderID,
		TakeOrderID: takeOrderID,
		Status:      PositionStatus(p.Status),
	}

	result := r.db.Create(positionDB)
	if result.Error != nil {
		return nil, result.Error
	}
	return positionDB.ToDomain(), nil
}

func (r *repository) Get(positionID string) *position.Position {
	positionDB := &Position{}
	res := r.db.First(positionDB, positionID)
	if res.Error != nil {
		return nil
	}
	return positionDB.ToDomain()
}

func (r *repository) Update(positionID string, fields map[string]interface{}) error {
	return r.db.Model(&Position{}).Where("position_id = ?", positionID).Updates(fields).Error
}
