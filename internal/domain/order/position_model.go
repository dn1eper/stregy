package order

type Position struct {
	ID        int64
	Size      float64
	MainOrder *Order
	CtgOrders []*Order
}

type PositionStatus int

const (
	DraftPosition PositionStatus = iota
	OpenPosition
	ClosedPosition
)

func (p *Position) Status() PositionStatus {
	if p.MainOrder.Status != FilledOrder {
		return DraftPosition
	}
	if p.Size == 0 {
		return ClosedPosition
	}
	return OpenPosition
}

func (p *Position) Copy() *Position {
	return &Position{
		ID:        p.ID,
		MainOrder: p.MainOrder,
		CtgOrders: p.CtgOrders,
	}
}

func (p *Position) AddCgtOrder(o *Order) {
	p.CtgOrders = append(p.CtgOrders, o)
}

func (p *Position) RemoveCgtOrder(id int64) {
	n := len(p.CtgOrders)
	for i, o := range p.CtgOrders {
		if o.ID == id {
			p.CtgOrders[i], p.CtgOrders[n-1] = p.CtgOrders[n-1], p.CtgOrders[i]
			p.CtgOrders = p.CtgOrders[:n-1]
			return
		}
	}
}
