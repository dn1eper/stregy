package posmerge

import (
	"fmt"
	"stregy/internal/domain/broker"
	"stregy/internal/domain/order"
)

type PositionMerger interface {
	MergePositions(positions ...*order.Position) (MergedPosition, error) // returns virtual position
	MergePosition(p *order.Position, mergedPositionID int64) (MergedPosition, error)
	SetStopLoss(id int64, price float64) error
	SetTakeProfit(id int64, price float64) error
	OnOrder(o order.Order)
	SetCallbackObject(co CallbackObject)
}

type CallbackObject interface {
	OnMergedPositionOrder(mp MergedPosition)
}

type positionMerger struct {
	broker         broker.Broker
	callbackObject CallbackObject

	mergedPositionCount int64
	mergedPositions     map[int64]*MergedPosition
	trackedOrders       map[int64]*MergedPosition
}

func NewPositionMerger(broker broker.Broker) PositionMerger {
	return &positionMerger{
		broker:          broker,
		mergedPositions: make(map[int64]*MergedPosition),
		trackedOrders:   make(map[int64]*MergedPosition)}
}

// returns: virtual position
func (pm *positionMerger) MergePositions(positions ...*order.Position) (MergedPosition, error) {
	if len(positions) == 0 {
		return MergedPosition{}, fmt.Errorf("no positions provided")
	}

	size := 0.0
	var slSize, tpSize float64
	diraction := positions[0].MainOrder.Diraction
	mainOrders := make(map[int64]*order.Order, 0)
	slOrders := make(map[int64]*order.Order, 0)
	tpOrders := make(map[int64]*order.Order, 0)
	for _, p := range positions {
		if err := pm.checkIsValidPosition(p); err != nil {
			return MergedPosition{}, err
		}
		if p.MainOrder.Diraction != diraction {
			return MergedPosition{}, fmt.Errorf("positions with different diractions cannot be merged")
		}

		size += p.Size
		mainOrders[p.MainOrder.ID] = p.MainOrder
		posSlSize, posTpSize := parseCtgOrders(p, slOrders, tpOrders)
		slSize += posSlSize
		tpSize += posTpSize
	}

	mergedPosition := MergedPosition{
		ID:         pm.mergedPositionCount,
		Size:       size,
		SlSize:     slSize,
		TpSize:     tpSize,
		Diraction:  diraction,
		mainOrders: mainOrders,
		slOrders:   slOrders,
		tpOrders:   tpOrders,
	}

	pm.mergedPositionCount += 1
	pm.mergedPositions[mergedPosition.ID] = &mergedPosition
	for _, o := range mergedPosition.slOrders {
		pm.trackedOrders[o.ID] = &mergedPosition
	}
	for _, o := range mergedPosition.tpOrders {
		pm.trackedOrders[o.ID] = &mergedPosition
	}

	return mergedPosition, nil
}

func (pm *positionMerger) MergePosition(p *order.Position, mergedPositionID int64) (MergedPosition, error) {
	if err := pm.checkIsValidPosition(p); err != nil {
		return MergedPosition{}, err
	}

	mergedPosition, ok := pm.mergedPositions[mergedPositionID]
	if !ok {
		return MergedPosition{}, fmt.Errorf("merged position not found")
	}

	mergedPosition.Size += p.Size
	posSlSize, posTpSize := parseCtgOrders(p, mergedPosition.slOrders, mergedPosition.tpOrders)
	mergedPosition.SlSize += posSlSize
	mergedPosition.TpSize += posTpSize

	mergedPosition.mainOrders[p.MainOrder.ID] = p.MainOrder
	for _, o := range mergedPosition.slOrders {
		pm.trackedOrders[o.ID] = mergedPosition
	}
	for _, o := range mergedPosition.tpOrders {
		pm.trackedOrders[o.ID] = mergedPosition
	}

	return *mergedPosition, nil
}

func (pm *positionMerger) SetStopLoss(id int64, price float64) error {
	mergedPosition, err := pm.getMergedPosition(id)
	if err != nil {
		return err
	}

	mainOrderIdToSlSizeToSubmit := make(map[int64]float64, len(mergedPosition.mainOrders))
	for _, p := range mergedPosition.mainOrders {
		mainOrderIdToSlSizeToSubmit[p.ID] = p.Size
	}
	for _, slOrder := range mergedPosition.slOrders {
		mainOrderIdToSlSizeToSubmit[slOrder.Position.MainOrder.ID] -= slOrder.Size

		if slOrder.Price == price {
			continue
		}

		if err := pm.broker.ChangeOrderPrice(slOrder.ID, price); err != nil {
			return err
		}
	}

	for id, size := range mainOrderIdToSlSizeToSubmit {
		if size <= 0 {
			continue
		}

		p := mergedPosition.mainOrders[id].Position
		slOrder, err := pm.broker.AddCtgOrder(
			p.ID,
			order.Order{
				Price:     price,
				Size:      size,
				Diraction: mergedPosition.Diraction.Opposite(),
				Type:      order.StopMarket})
		if err != nil {
			return err
		}
		mergedPosition.slOrders[slOrder.ID] = slOrder
	}

	return nil
}

func (pm *positionMerger) SetTakeProfit(id int64, price float64) error {
	mergedPosition, err := pm.getMergedPosition(id)
	if err != nil {
		return err
	}

	mainOrderIdToTpSizeToSubmit := make(map[int64]float64, len(mergedPosition.mainOrders))
	for _, p := range mergedPosition.mainOrders {
		mainOrderIdToTpSizeToSubmit[p.ID] = p.Size
	}
	for _, tpOrder := range mergedPosition.tpOrders {
		mainOrderIdToTpSizeToSubmit[tpOrder.Position.MainOrder.ID] -= tpOrder.Size

		if tpOrder.Price == price {
			continue
		}

		err := pm.broker.ChangeOrderPrice(tpOrder.ID, price)
		if err != nil {
			return err
		}
	}

	for id, size := range mainOrderIdToTpSizeToSubmit {
		if size <= 0 {
			continue
		}

		p := mergedPosition.mainOrders[id].Position
		tpOrder, err := pm.broker.AddCtgOrder(
			p.ID,
			order.Order{
				Price:     price,
				Size:      size,
				Diraction: mergedPosition.Diraction.Opposite(),
				Type:      order.Limit})
		if err != nil {
			return err
		}
		mergedPosition.tpOrders[tpOrder.ID] = tpOrder
	}

	return nil
}

func (pm *positionMerger) OnOrder(o order.Order) {
	mergedPosition, ok := pm.trackedOrders[o.ID]
	if !ok {
		return
	}

	_, isSlOrder := mergedPosition.slOrders[o.ID]
	_, isTpOrder := mergedPosition.tpOrders[o.ID]
	if !isSlOrder && !isTpOrder { // order can't be in mainOrders because only Filled main orders are accepted
		panic("invariant violation: order is tracked but not present in it's MergedPosition")
	}

	switch o.Status {
	case order.FilledOrder:
		if isSlOrder {
			mergedPosition.ClosedBySL += o.Size
			mergedPosition.SlSize -= o.Size
			delete(mergedPosition.slOrders, o.ID)
		} else {
			mergedPosition.ClosedByTP += o.Size
			mergedPosition.TpSize -= o.Size
			delete(mergedPosition.tpOrders, o.ID)
		}
		mergedPosition.Size -= o.Size

		if mergedPosition.Size < 0 {
			panic("merged position got subzero size")
		}
		if mergedPosition.Size == 0 {
			pm.deleteMergedPosition(mergedPosition.ID)
		}

	case order.CancelledOrder:
		if isSlOrder {
			mergedPosition.SlSize -= o.Size
		} else {
			mergedPosition.TpSize -= o.Size
		}
		pm.removeCtgOrder(o.ID)
	}

	if pm.callbackObject != nil {
		pm.callbackObject.OnMergedPositionOrder(*mergedPosition)
	}
}

func (pm *positionMerger) SetCallbackObject(co CallbackObject) {
	pm.callbackObject = co
}
