package strat1

import (
	"math"
	"stregy/internal/domain/order"
)

/*
	Idea: enter when 3 consequent bars closed in the same diraction,
	avarage when price goes 1 tp to the opposite diraction, use fixed ammount sl
*/

const startSize = 1.0
const tpPrice = 200.0
const slAmount = 1000.0
const newSLMinDistanceFromPrice = 50.0

const conseqClosesToTrade int = 3

func (s *strat1) buy() (err error) {
	slPrice := s.broker.Price() - (slAmount / startSize)
	tpPrice := s.broker.Price() + tpPrice

	newMainOrder, err := s.broker.SubmitOrder(
		order.Order{
			Size:      startSize,
			Diraction: order.Long,
			Type:      order.Market},
		order.Order{
			Price:     slPrice,
			Size:      startSize,
			Diraction: order.Short,
			Type:      order.StopMarket,
		},
		order.Order{
			Price:     tpPrice,
			Size:      startSize,
			Diraction: order.Short,
			Type:      order.Limit,
		},
	)
	if err == nil {
		s.newMainOrder = newMainOrder
	}

	return err
}

func (s *strat1) sell() (err error) {
	slPrice := s.broker.Price() + (slAmount / startSize)
	tpPrice := s.broker.Price() - tpPrice

	newMainOrder, err := s.broker.SubmitOrder(
		order.Order{
			Size:      startSize,
			Diraction: order.Short,
			Type:      order.Market},
		order.Order{
			Price:     slPrice,
			Size:      startSize,
			Diraction: order.Long,
			Type:      order.StopMarket,
		},
		order.Order{
			Price:     tpPrice,
			Size:      startSize,
			Diraction: order.Long,
			Type:      order.Limit,
		},
	)
	if err == nil {
		s.newMainOrder = newMainOrder
	}

	return err
}

// double total position size, set tp to the mean position price,
// if resulting sl closer than 50 to price then do nothing
func (s *strat1) avarage() error {
	mergedPosition := s.mergedPosition
	diraction := mergedPosition.Diraction
	size := mergedPosition.Size
	newSize := size * 2
	newMeanPrice := (s.mergedPosition.MeanPrice()*size + (s.broker.Price() * size)) / newSize
	newTpPrice := newMeanPrice
	var newSLPrice float64
	if diraction == order.Long {
		newSLPrice = newMeanPrice - (slAmount / newSize) // slAmount = (price - slPrice) * size -> slPrice = price - slAmount / size
	} else {
		newSLPrice = newMeanPrice + (slAmount / newSize) // slAmount = (slPrice - price) * size -> slPrice = slAmount / size - price
	}

	if math.Abs(newSLPrice-s.broker.Price()) < newSLMinDistanceFromPrice {
		return nil
	}

	// place new order
	o, err := s.broker.SubmitOrder(
		order.Order{
			Size:      size,
			Diraction: diraction,
			Type:      order.Market},
		order.Order{
			Price:     newSLPrice,
			Size:      size,
			Diraction: diraction.Opposite(),
			Type:      order.StopMarket},
		order.Order{
			Price:     newTpPrice,
			Size:      size,
			Diraction: diraction.Opposite(),
			Type:      order.Limit},
	)
	if err != nil {
		return err
	}
	s.newMainOrder = o
	s.newSlPrice = newSLPrice
	s.newTpPrice = newTpPrice

	return nil
}
