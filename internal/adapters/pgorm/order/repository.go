package order

import (
	"stregy/internal/domain/order"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(client *gorm.DB) order.Repository {
	return &repository{db: client}
}

func (r *repository) Create(o *order.Order, stratexecID string) (*order.Order, error) {
	stratexecIDParsed, _ := uuid.Parse(stratexecID)

	orderDB := &Order{
		StrategyExecutionID: stratexecIDParsed,
		Price:               o.Price,
		Direction:           OrderDirection(o.Direction),
		Size:                o.Size,
		Type:                OrderType(o.Type),
		SetupTime:           o.SetupTime,
		DoneTime:            o.DoneTime,
		FillPrice:           o.FillPrice,
		Status:              OrderStatus(o.Status),
	}

	result := r.db.Create(orderDB)
	if result.Error != nil {
		return nil, result.Error
	}
	return orderDB.ToDomain(), nil
}

func (r *repository) Get(orderID string) *order.Order {
	orderDB := &Order{}
	parsedOrderID, _ := uuid.Parse(orderID)
	res := r.db.First(orderDB, parsedOrderID)
	if res.Error != nil {
		return nil
	}
	return orderDB.ToDomain()
}

func (r *repository) Update(orderID string, fields map[string]interface{}) error {
	return r.db.Model(&Order{}).Where("order_id = ?", orderID).Updates(fields).Error
}
