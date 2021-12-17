package ch01

import (
	"strings"
)

func uRLIfyString(text string) string {
	text = strings.TrimSpace(text)
	text = strings.Replace(text, " ", "%20", -1)
	return text
}
