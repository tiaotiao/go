package goutil

import (
	"strings"
)

// Captialize the first letter
func ToProper(s string) string {
	if s == "" {
		return ""
	}
	a := strings.ToUpper(s[:1])
	b := strings.ToLower(s[1:])
	return a + b
}
