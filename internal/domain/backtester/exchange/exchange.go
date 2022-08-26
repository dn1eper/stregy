package exchange

import (
	"stregy/internal/domain/exchange"
	"stregy/internal/domain/order"
	"stregy/internal/domain/position"
	"stregy/internal/domain/quote"
	"stregy/internal/domain/tick"
	"time"

	"github.com/google/uuid"
	btree "github.com/ross-oreto/go-tree"
	log "github.com/sirupsen/logrus"
)

type simulatedExchange struct {
	broker   exchange.Broker
	lastTick tick.Tick

	timeframe  int
	barsNeeded int

	tickFeed  <-chan tick.Tick
	quoteFeed <-chan quote.Quote

	activeOrders       *btree.Btree
	activeMarketOrders *btree.Btree
	activePositions    *btree.Btree
}

func NewExchange(b exchange.Broker, tickFeed <-chan tick.Tick, quoteFeed <-chan quote.Quote, timeframe, barsNeeded int) exchange.Exchange {
	return &simulatedExchange{
		timeframe: timeframe,
		broker:    b,
		tickFeed:  tickFeed,
		quoteFeed: quoteFeed,

		activeOrders:       btree.New(),
		activeMarketOrders: btree.New(),
		activePositions:    btree.New(),
	}
}

// if skipBars != 0 calls FirstBars with []quote.Quote
func (e *simulatedExchange) Run() error {
	tickFeed := e.tickFeed
	quoteFeed := e.quoteFeed

	tfDuration := time.Minute * time.Duration(e.timeframe)

	var firstBarTime time.Time
	if e.barsNeeded > 0 {
		quotes := make([]quote.Quote, 0, e.barsNeeded)
		for i := 0; i < e.barsNeeded; i++ {
			quotes = append(quotes, <-e.quoteFeed)
		}
		e.broker.FirstBars(quotes)
		firstBarTime = quotes[len(quotes)-1].Time.Add(tfDuration)
	}

	// pull unnecessary ticks
	t := <-tickFeed
	for t.Time.Before(firstBarTime) {
		t = <-tickFeed
	}
	nextBarStart := nextBarStart(t.Time, e.timeframe)
	// fmt.Printf("next bar start: %v\n", nextBarStart)
	for t := range tickFeed {
		e.lastTick = t

		if !t.Time.Before(nextBarStart) {
			nextBarStart = nextBarStart.Add(tfDuration)

			q, ok := <-quoteFeed
			if !ok {
				log.Info("Quote channel has been closed.")
				break
			}
			e.broker.OnQuote(&q)
		}

		// exchange logic
		for _, orderI := range e.activeMarketOrders.Values() {
			o := orderI.(*Order)
			if o.Type == order.Market {
				e.openPosition(o, e.price())
			} else if o.Type == order.CloseByMarket {
				e.closePosition(o, e.price())
			} else {
				panic("incorrect order type in activeMarketOrders")
			}
		}

		for _, orderI := range e.activeOrders.Values() {
			o := orderI.(*Order)

			if o.Type == order.TrailingStop || o.Type == order.StopLimit {
				panic("not implemented")
			}

			if (o.Direction == order.Long && o.abovePrice && t.Price >= o.Price) ||
				(o.Direction == order.Long && !o.abovePrice && t.Price <= o.Price) ||
				(o.Direction != order.Long && o.abovePrice && t.Price <= o.Price) ||
				(o.Direction != order.Long && !o.abovePrice && t.Price >= o.Price) {
				e.trade(o)
			}
			// if o.Direction == order.Long {
			// 	if o.abovePrice {
			// 		if t.Price >= o.Price {
			// 			e.trade(o)
			// 		}
			// 	} else {
			// 		if t.Price <= o.Price {
			// 			e.trade(o)
			// 		}
			// 	}
			// } else {
			// 	if o.abovePrice {
			// 		if t.Price <= o.Price {
			// 			e.trade(o)
			// 		}
			// 	} else {
			// 		if t.Price >= o.Price {
			// 			e.trade(o)
			// 		}
			// 	}
			// }
		}
	}

	return nil
}

func (e *simulatedExchange) trade(o *Order) {
	switch o.Type {
	case order.CloseByLimit:
		e.closePosition(o, o.Price)
	case order.CloseByStop:
		e.closePosition(o, e.price())
	case order.Limit:
		e.openPosition(o, o.Price)
	case order.Stop:
		e.openPosition(o, e.price())
	}
}

func (e *simulatedExchange) RegisterPosition(p *position.Position) {
	p.MainOrder.SetupTime = e.lastTick.Time
	if p.MainOrder.Type != order.Market {
		e.addActiveMarketOrder(&p.MainOrder)
	} else {
		e.addActiveOrder(&p.MainOrder)
	}
	e.activePositions.Insert(&Position{p})
}

func (e *simulatedExchange) CancelOrder(o *order.Order) {
	o.DoneTime = e.lastTick.Time
	e.activeOrders.Delete(&Order{o, false})
	e.activePositions.Delete(&Position{&position.Position{PositionID: o.PositionID}})
}

func (e *simulatedExchange) ClosePosition(p *position.Position) {
	if e.activePositions.Contains(&Position{p}) {
		closeByMarketOrder := order.Order{
			OrderID:    uuid.New().String(),
			Direction:  order.OppositeDirection(p.MainOrder.Direction),
			Size:       p.MainOrder.Size,
			Type:       order.CloseByMarket,
			PositionID: p.PositionID,
		}
		e.addActiveMarketOrder(&closeByMarketOrder)
	}
}

func (e *simulatedExchange) openPosition(o *Order, fillPrice float64) {
	e.activeOrders.Delete(o)
	o.DoneTime = e.lastTick.Time
	o.FillPrice = fillPrice
	o.Status = order.Filled

	p := e.getPosition(o)
	if p.StopOrder != nil {
		p.StopOrder.SetupTime = e.lastTick.Time
		e.addActiveOrder(p.StopOrder)
	}
	if p.TakeOrder != nil {
		p.TakeOrder.SetupTime = e.lastTick.Time
		e.addActiveOrder(p.TakeOrder)
	}
}

func (e *simulatedExchange) closePosition(o *Order, fillPrice float64) {
	p := e.getPosition(o)
	cancelOrders := make([]*order.Order, 0, 2)
	if p.StopOrder != nil && o.OrderID == p.StopOrder.OrderID {
		p.Status = position.StopLoss
		cancelOrders = append(cancelOrders, p.TakeOrder)
	} else if p.TakeOrder != nil && o.OrderID == p.TakeOrder.OrderID {
		p.Status = position.TakeProfit
		cancelOrders = append(cancelOrders, p.StopOrder)
	} else {
		p.Status = position.MarketClose
		cancelOrders = append(cancelOrders, p.StopOrder, p.TakeOrder)
	}

	for _, oCancel := range cancelOrders {
		if oCancel != nil {
			e.finalizeOrder(oCancel, order.Cancelled)
		}
	}
	o.FillPrice = fillPrice
	e.finalizeOrder(o.Order, order.Filled)
}

func (e *simulatedExchange) time() time.Time {
	return e.lastTick.Time
}

func (e *simulatedExchange) price() float64 {
	return e.lastTick.Price
}

func (e *simulatedExchange) getPosition(o *Order) *Position {
	return e.activePositions.Get(&Position{&position.Position{PositionID: o.PositionID}}).(*Position)
}

func (e *simulatedExchange) addActiveOrder(p *order.Order) {
	e.activeOrders.Insert(&Order{p, p.Price > e.lastTick.Price})
}

func (e *simulatedExchange) addActiveMarketOrder(p *order.Order) {
	e.activeMarketOrders.Insert(&Order{p, p.Price > e.lastTick.Price})
}

func (e *simulatedExchange) deleteActiveOrder(o *order.Order) {
	e.activeOrders.Delete(&Order{o, false})
}

func (e *simulatedExchange) finalizeOrder(o *order.Order, status order.OrderStatus) {
	o.Status = status
	o.DoneTime = e.time()
	e.deleteActiveOrder(o)
}

// if t is alligned with timeframe returns bar started at t, else (t + timeframe)
func nextBarStart(t time.Time, timeframe int) time.Time {
	curBarT := t.Truncate(time.Duration(timeframe) * time.Minute)
	if curBarT == t {
		return t
	}
	return curBarT.Add(time.Duration(timeframe) * time.Minute).Truncate(time.Minute)
}

// For broker

// func TR(prevQuote, curQuote quote.Quote) float64 {
// 	max := curQuote.High - curQuote.Low
// 	max = math.Max(max, math.Abs(curQuote.High-prevQuote.Close))
// 	max = math.Max(max, math.Abs(curQuote.Low-prevQuote.Close))
// 	return max
// }

// func UpdateATRincremental() {
// 	if bt.ATRperiod != 0 {
// 		ATR = ((ATR * float64(bt.ATRperiod-1)) + (High(0) - Low(0))) / float64(bt.ATRperiod)
// 	}
// }
