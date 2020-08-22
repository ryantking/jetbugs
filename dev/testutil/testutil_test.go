package testutil_test

// INTERNAL

import (
	"testing"

	"github.com/ryantking/jetbugs/dev/testutil"
)

func TestAssertEqual(t *testing.T) {
	tests := []struct{
		name string
		x, y interface{}
		equal bool
	}{
		{"Equal", 1, 1, true},
		{"NotEqual", 1, 2, false},
		{"SameValueDifferentTypes", 1, 1.0, false},
	}

	for _, tt := range tests {
		testT := new(testing.T)
		equal := testutil.AssertEqual(testT, tt.x, tt.y)
		if equal != tt.equal {
			t.Fatal("can't even write an assert function")
		}
		if tt.equal && testT.Failed() {
			t.Fatal("equal values, but test still failed")
		} else if !tt.equal && !testT.Failed() {
			t.Fatal("not equal values, but test didn't fail")
		}
	}
}

