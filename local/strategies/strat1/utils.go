package strat1

import (
	"stregy/internal/domain/order"
	"stregy/internal/domain/quote"
)

func (s *strat1) havePosition() bool {
	return s.mergedPosition != nil
}

func getNextAveragingPrice(diraction order.OrderDiraction, startPrice float64) float64 {
	if diraction == order.Long {
		return startPrice - tpPrice
	} else {
		return startPrice + tpPrice
	}
}

func (s *strat1) updatePrevConseqClosesCount(q quote.Quote) {
	if s.prevClose == 0 {
		s.prevClose = q.Close
	} else if q.Close > s.prevClose {
		s.prevClosesUp += 1
		s.prevClosesDown = 0
	} else if q.Close < s.prevClose {
		s.prevClosesDown += 1
		s.prevClosesUp = 0
	}
}

func (s *strat1) averageIfNeeded(price float64) {
	var err error

	if s.mergedPosition.Diraction == order.Long {
		if price <= s.nextAveragingPrice {
			err = s.avarage()
		}
	} else {
		if price >= s.nextAveragingPrice {
			err = s.avarage()
		}
	}

	if err != nil {
		panic(err)
	}
}

func (s *strat1) openPositionIfNeeded() (err error) {
	if s.prevClosesUp >= conseqClosesToTrade {
		err = s.buy()
	} else if s.prevClosesDown >= conseqClosesToTrade {
		err = s.sell()
	}

	return err
}
