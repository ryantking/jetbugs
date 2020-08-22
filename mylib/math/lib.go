package math

import "math"

// EXTERNAL

// Pow ....
func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
