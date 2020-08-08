package mylib_test

import (
	"testing"

	"github.com/ryantking/jetbugs/mylib"
)

func TestSquare(t *testing.T) {
	res := mylib.Pow(2, 3)
	if res != 8 {
		t.Fatalf("wrong result: %d, expected 8", res)
	}
}

func TestSquare2(t *testing.T) {
	t.Log("this test won't run")
}