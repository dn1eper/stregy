package utils

import "time"

func FormatTime(t time.Time) string {
	return t.Format("2006-01-02 03:04:05.000 -0700")
}
