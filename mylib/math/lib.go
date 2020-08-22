package math

import "math"

// EXTERNAL

// Pow ....
func Pow(x, y int) int {
	// only added to not get 100% coverage
	if y == 0 {
		return 1
	}

	return int(math.Pow(float64(x), float64(y)))
}
