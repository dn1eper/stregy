package bt

import "fmt"

type OrderNotFoundError struct {
	OrderID int64
}

func (err *OrderNotFoundError) Error() string {
	return fmt.Sprintf("order %d not found", err.OrderID)
}

type InvalidOrderError struct {
	OrderID int64
	Err     error
}

func (err *InvalidOrderError) Error() string {
	return fmt.Sprintf("order %d is invalid: %s", err.OrderID, err.Err.Error())
}

type PositionNotFoundError struct {
	PositionID int64
}

func (err *PositionNotFoundError) Error() string {
	return fmt.Sprintf("position %d not found", err.PositionID)
}
