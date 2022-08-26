package strategy

import (
	"fmt"
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/strategy"
	. "stregy/internal/domain/strategy/core"

	log "github.com/sirupsen/logrus"
)

type Strategy struct{}

func NewStrategy() strategy.Implementation {
	log.Info(fmt.Sprintf("%v: starting strategy\n", Time(0)))
	return &Strategy{}
}

func (s *Strategy) Config() strategy.StrategyConfig {
	return strategy.StrategyConfig{
		BarsNeeded: 40,
		ATRperiod:  21,
	}
}

func (s *Strategy) OnQuote(quote quote.Quote) {

}

func (s *Strategy) OnOrder(order order.Order) {

}

func (s *Strategy) OnPosition(position position.Position) {

}

func (s *Strategy) OnExit() {
	log.Info(fmt.Sprintf("%v: quiting strategy\n", Time(0)))
}
