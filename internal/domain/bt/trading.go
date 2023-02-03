package bt

import (
	"fmt"
	"stregy/internal/domain/order"
)

var orders map[int64]*order.Order
var positions map[int64]*order.Position
var orderCount int64
var positionCount int64

var orderHistory []*order.Order

func SubmitOrder(price float64, size float64, diraction order.OrderDiraction, orderType order.OrderType) error {
	o := &order.Order{
		Price:     price,
		Size:      size,
		Diraction: diraction,
		Type:      orderType,
	}
	p := &order.Position{MainOrder: o, Status: order.CreatedPosition}
	o.Position = p

	return submitOrder(o)
}

func submitOrder(o *order.Order) error {
	if o == nil {
		return nil
	}

	err := checkIsValidOrder(o)
	if err != nil {
		return err
	}

	o.ID = orderCount
	o.SubmissionTime = Time
	o.Status = order.Submitted
	orders[o.ID] = o
	orderCount += 1

	orderHistory = append(orderHistory, o)

	logOrderStatusChange(o)
	strategy.OnOrder(*o)

	return nil
}

func CancelOrder(id int64) error {
	o, ok := orders[id]
	if !ok {
		return fmt.Errorf("order not found")
	}

	delete(orders, id)

	o.Status = order.CancelledOrder
	o.FCTime = Time

	logOrderStatusChange(o)
	strategy.OnOrder(*o)

	return nil
}

func SubmitContingentOrders(price float64, size float64, diraction order.OrderDiraction, orderType order.OrderType, sl, tp float64) (*order.Position, error) {
	o := &order.Order{
		Price:          price,
		Size:           size,
		Diraction:      diraction,
		Type:           orderType,
		SubmissionTime: Time,
	}

	err := checkIsValidOrder(o)
	if sl <= 0 || tp <= 0 {
		err = fmt.Errorf("sl and tp must be greater than 0")
	}
	if err != nil {
		return nil, err
	}

	var slPrice, tpPrice float64
	if o.Diraction == order.Long {
		slPrice = o.Price - sl
		tpPrice = o.Price + tp
	} else {
		slPrice = o.Price + sl
		tpPrice = o.Price - tp
	}

	slOrder := &order.Order{
		Diraction: o.Diraction.Opposite(),
		Size:      o.Size,
		Price:     slPrice,
		Status:    order.Inactive,
		Type:      order.StopMarket,
	}
	tpOrder := &order.Order{
		Diraction: o.Diraction.Opposite(),
		Size:      o.Size,
		Price:     tpPrice,
		Status:    order.Inactive,
		Type:      order.Limit,
	}
	position := &order.Position{
		MainOrder: o,
		StopOrder: slOrder,
		TakeOrder: tpOrder,
		Status:    order.CreatedPosition,
	}
	o.Position = position
	slOrder.Position = position
	tpOrder.Position = position
	submitOrder(o)

	return position, nil
}

func executeOrder(o *order.Order, price float64) {
	o.Status = order.Filled
	o.ExecutionPrice = price
	o.FCTime = Time
	delete(orders, o.ID)

	logOrderStatusChange(o)
	strategy.OnOrder(*o)

	handleContingentOrders(o)
}

func handleContingentOrders(o *order.Order) {
	if o.ID == o.Position.MainOrder.ID {
		o.Position.ID = positionCount
		positionCount += 1
		o.Position.Status = order.OpenPosition
		submitOrder(o.Position.StopOrder)
		submitOrder(o.Position.TakeOrder)

	} else if o.ID == o.Position.StopOrder.ID {
		o.Position.Status = order.StopLoss
		if o.Position.TakeOrder != nil {
			CancelOrder(o.Position.TakeOrder.ID)
		}

	} else {
		o.Position.Status = order.TakeProfit
		if o.Position.StopOrder != nil {
			CancelOrder(o.Position.StopOrder.ID)
		}
	}
}

func checkIsValidOrder(o *order.Order) (err error) {
	if o.Size <= 0 {
		err = fmt.Errorf("invalid order size %v", o.Size)
	}
	if o.Price <= 0 {
		err = fmt.Errorf("invalid order price %v", o.Price)
	}

	return err
}
