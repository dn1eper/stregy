package exchange

import (
	"stregy/internal/domain/position"

	btree "github.com/ross-oreto/go-tree"
)

type Position struct {
	*position.Position
}

func (o *Position) Comp(than btree.Val) int8 {
	if o.PositionID < than.(*Position).PositionID {
		return -1
	} else if o.PositionID > than.(*Position).PositionID {
		return 1
	}
	return 0
}
