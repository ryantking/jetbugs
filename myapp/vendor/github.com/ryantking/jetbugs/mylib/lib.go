package mylib

import "math"

// Pow ....
func Pow(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
