package gophp

import (
	"errors"
	"testing"
)

func TestMd5(t *testing.T) {
	dst := Md5("apple")
	if dst != "1f3870be274f6c49b3e31a0c6728957f" {
		t.Error(errors.New("the md5 hash of 'apple' should be 1f3870be274f6c49b3e31a0c6728957f"))
	}
}
