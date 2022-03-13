package library

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	oldZ := 0.0
	// choose arbitrary value depending on your "fault tolerance"
	for math.Abs(z-oldZ) > 0.000000000001 {
		oldZ = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func RunSqrt() {
	fmt.Println(Sqrt(2))
}
