package strat1

import "stregy/local/strategies/posmerge"

func (s *strat1) OnMergedPositionOrder(mp posmerge.MergedPosition) {
	s.mergedPosition = &mp
	if mp.Size == 0 {
		s.mergedPosition = nil
	}
}
