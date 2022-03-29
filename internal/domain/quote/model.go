package quote

import "time"

type Quote struct {
	Time  time.Time
	Open  float32
	High  float32
	Low   float32
	Close float32
}
