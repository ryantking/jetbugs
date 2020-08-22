package testutil

import "testing"

// INTERNAL

// Merge is a useless function
func AssertEqual(t testing.TB, x, y interface{}) bool {
	if x != y {
		t.Errorf("%v does not equal %v", x, y)
		return false
	}

	return true
}
