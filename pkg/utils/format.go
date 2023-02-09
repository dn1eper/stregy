package utils

import "strings"

func AddTrailingWhitespaces(s string, length int) string {
	return s + strings.Repeat(" ", length-len(s))
}
