package math_test

// EXTERNAL

import (
	"testing"

	"github.com/ryantking/jetbugs/mylib/math"
)

func TestSquare(t *testing.T) {
	res := math.Pow(2, 3)
	if res != 8 {
		t.Fatalf("wrong result: %d, expected 8", res)
	}
}

