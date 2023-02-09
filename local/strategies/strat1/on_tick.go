package strat1

func (s *strat1) OnTick(price float64) {
	if s.havePosition() {
		s.averageIfNeeded(price)
	}
}
