package zerotriangles

import (
	"math"
)

// Function to calculate the circumradius of an equilateral triangle given the side length
func CircumradiusEquilateralTriangle(a float64) float64 {
	// Circumradius formula for an equilateral triangle
	R := a / math.Sqrt(3)
	return R
}
