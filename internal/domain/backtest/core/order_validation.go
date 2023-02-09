package core

import (
	"fmt"
	"stregy/internal/domain/order"
)

func checkIsValidOrder(o *order.Order) (err error) {
	if o.Size <= 0 {
		return &InvalidOrderError{o.ID, fmt.Errorf("size is less than zero")}
	}
	if o.Price <= 0 && o.Type != order.Market {
		return &InvalidOrderError{o.ID, fmt.Errorf("price is less than zero")}
	}
	err = checkIsSupportedOrderType(o)

	return err
}

func checkIsValidCtgOrder(o, mainOrder *order.Order) (err error) {
	err = checkIsValidOrder(o)
	if err != nil {
		return err
	}

	if o.Size > mainOrder.Size {
		return &InvalidOrderError{o.ID, fmt.Errorf("contingent order size is greater than main order size")}
	}
	if o.Diraction != mainOrder.Diraction.Opposite() {
		return &InvalidOrderError{o.ID, fmt.Errorf("contingent order diraction is not opposite")}
	}

	return nil
}

func checkIsSupportedOrderType(o *order.Order) (err error) {
	if !(o.Type == order.Limit || o.Type == order.Market || o.Type == order.StopMarket) {
		return fmt.Errorf("order type %s is not supported", o.Type.String())
	}

	return nil
}
