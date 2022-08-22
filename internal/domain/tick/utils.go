package tick

import (
	"context"
	"time"
)

func tickGenerator(ctx context.Context, ch chan<- Tick, s service, symbol string, start, end time.Time) {
	batchStart := start
	batchEnd := batchStart.AddDate(0, 0, 1)
	if batchEnd.After(end) {
		batchEnd = end
	}

	for true {
		ticks, err := s.repository.GetByInterval(ctx, symbol, batchStart, batchEnd)
		if err != nil {
			panic(err)
		}
		if len(ticks) == 0 {
			break
		}

		for _, tick := range ticks {
			ch <- tick
		}

		batchStart = batchEnd
		batchEnd = batchStart.AddDate(0, 0, 1)
		if batchEnd.After(end) {
			batchEnd = end
		}
	}
	close(ch)
}
