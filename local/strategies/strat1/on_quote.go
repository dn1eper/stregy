package strat1

import (
	"stregy/internal/domain/quote"
)

func (s *strat1) OnQuote(q quote.Quote, timeframe int) {
	var err error

	s.updatePrevConseqClosesCount(q)

	if !s.havePosition() {
		err = s.openPositionIfNeeded()
	}

	if err != nil {
		panic(err)
	}

	s.prevClose = q.Close
}
