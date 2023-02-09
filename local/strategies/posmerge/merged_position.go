package posmerge

import "stregy/internal/domain/order"

type MergedPosition struct {
	ID         int64
	Size       float64
	Diraction  order.OrderDiraction
	SlSize     float64
	TpSize     float64
	ClosedBySL float64
	ClosedByTP float64
	mainOrders map[int64]*order.Order
	slOrders   map[int64]*order.Order
	tpOrders   map[int64]*order.Order
}

func (pm *MergedPosition) ContainsOrder(id int64) bool {
	_, isMainOrder := pm.mainOrders[id]
	_, isSlOrder := pm.slOrders[id]
	_, isTpOrder := pm.tpOrders[id]
	return isMainOrder || isSlOrder || isTpOrder
}

func (pm *MergedPosition) MeanPrice() float64 {
	totalMoney := 0.0
	totalSize := 0.0
	for _, o := range pm.mainOrders {
		totalMoney += o.ExecutionPrice * o.Size
		totalSize += o.Size
	}

	return totalMoney / totalSize
}
