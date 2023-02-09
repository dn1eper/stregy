package posmerge

import (
	"fmt"
	"stregy/internal/domain/order"
)

func (pm *positionMerger) checkIsValidPosition(p *order.Position) error {
	if p.Status() != order.OpenPosition {
		return fmt.Errorf("cannot merge position that is not open")
	}

	if _, ok := pm.trackedOrders[p.MainOrder.ID]; ok {
		return fmt.Errorf("position %d already merged", p.ID)
	}

	return nil
}

func (pm *positionMerger) getMergedPosition(id int64) (*MergedPosition, error) {
	mergedPosition, ok := pm.mergedPositions[id]
	if !ok {
		return nil, fmt.Errorf("could not find virtual position %d", id)
	}

	return mergedPosition, nil
}

func (pm *positionMerger) removeCtgOrder(id int64) {
	mergedPosition, ok := pm.trackedOrders[id]
	if !ok {
		return
	}

	delete(mergedPosition.slOrders, id)
	delete(mergedPosition.tpOrders, id)
}

func (pm *positionMerger) deleteMergedPosition(id int64) {
	mergedPosition, ok := pm.mergedPositions[id]
	if !ok {
		return
	}

	delete(pm.mergedPositions, id)
	for id := range mergedPosition.slOrders {
		delete(pm.trackedOrders, id)
	}
	for id := range mergedPosition.tpOrders {
		delete(pm.trackedOrders, id)
	}
}

func parseCtgOrders(p *order.Position, destSlOrders, destTpOrders map[int64]*order.Order) (slSize float64, tpSize float64) {
	for _, ctgOrder := range p.CtgOrders {
		if ctgOrderIsSL(ctgOrder, p.MainOrder) {
			destSlOrders[ctgOrder.ID] = ctgOrder
			slSize += ctgOrder.Size
		} else {
			destTpOrders[ctgOrder.ID] = ctgOrder
			tpSize += ctgOrder.Size
		}
	}

	return slSize, tpSize
}

func ctgOrderIsSL(ctgOrder, mainOrder *order.Order) bool {
	x := mainOrder.ExecutionPrice - ctgOrder.Price
	if mainOrder.Diraction == order.Long {
		return x > 0
	} else {
		return x < 0
	}
}
