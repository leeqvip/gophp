package gophp

import (
	"strings"
)

// Strtolower Make a string lowercase
func Strtolower(str string) string {
	return strings.ToLower(str)
}

// Strtoupper Make a string uppercase
func Strtoupper(str string) string {
	return strings.ToUpper(str)
}
