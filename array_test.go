package gophp

import (
	"log"
	"testing"
)

func TestImplode(t *testing.T) {
	var array []string
	array = append(array, "a")
	array = append(array, "b")
	array = append(array, "c")

	r := Implode(", ", array)
	log.Println(r)
	if r != "a, b, c" {
		t.Error(r)
	}
}
