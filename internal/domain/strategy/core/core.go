package core

import (
	"stregy/internal/domain/account"
	"stregy/internal/domain/dataseries"
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/strategy"
	"time"
)

var broker strategy.Broker
var dataSeries dataseries.DataSeries
var acc account.Account

func SendOrder(direction order.OrderDirection, size float64, price float64, orderType order.OrderType) (*order.Order, error) {
	return broker.SendOrder(direction, size, price, orderType)
}
func CancelOrder(orderID string) error {
	return broker.CancelOrder(orderID)
}
func ClosePosition(positionID string) error {
	return broker.ClosePosition(positionID)
}

func Time(i int) time.Time {
	return dataSeries.Quotes[len(dataSeries.Quotes)-1-i].Time
}
func Open(i int) float64 {
	return dataSeries.Quotes[len(dataSeries.Quotes)-1-i].Open
}
func High(i int) float64 {
	return dataSeries.Quotes[len(dataSeries.Quotes)-1-i].High
}
func Low(i int) float64 {
	return dataSeries.Quotes[len(dataSeries.Quotes)-1-i].Low
}
func Close(i int) float64 {
	return dataSeries.Quotes[len(dataSeries.Quotes)-1-i].Close
}
func Volume(i int) float64 {
	return dataSeries.Quotes[len(dataSeries.Quotes)-1-i].Volume
}

func ActiveOrders() []order.Order {
	return acc.ActiveOrders()
}
func ActivePositions() []position.Position {
	return acc.ActivePositions()
}
