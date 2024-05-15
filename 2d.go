package zerotriangles

import "math"

// Function to triangulate an array of points in 2D
func TriangulateFan2D(a [2]float64, p [][2]float64) [][3][2]float64 {
	// Array to store the resulting triangles
	var triangles [][3][2]float64

	// Iterate over the points in p and form triangles with a
	for i := 0; i < len(p)-1; i++ {
		triangle := [3][2]float64{a, p[i], p[i+1]}
		triangles = append(triangles, triangle)
	}

	return triangles
}

// Function to calculate a circumscribing equilateral triangle around a circle in 2D
func TriangulateCircle2D(center [2]float64, radius float64) [3][2]float64 {
	// Side length of the equilateral triangle
	sideLength := radius * math.Sqrt(3)

	// Calculate the vertices of the equilateral triangle
	angle1 := math.Pi / 6          // 30 degrees
	angle2 := angle1 + 2*math.Pi/3 // 150 degrees
	angle3 := angle2 + 2*math.Pi/3 // 270 degrees

	v1 := [2]float64{
		center[0] + sideLength*math.Cos(angle1),
		center[1] + sideLength*math.Sin(angle1),
	}
	v2 := [2]float64{
		center[0] + sideLength*math.Cos(angle2),
		center[1] + sideLength*math.Sin(angle2),
	}
	v3 := [2]float64{
		center[0] + sideLength*math.Cos(angle3),
		center[1] + sideLength*math.Sin(angle3),
	}

	return [3][2]float64{v1, v2, v3}
}

// Function to triangulate a point in 2D space
func TriangulateAngle2D(a, b [2]float64, angle float64) [2]float64 {
	// Calculate the distance between points a and b
	r := math.Sqrt((b[0]-a[0])*(b[0]-a[0]) + (b[1]-a[1])*(b[1]-a[1]))

	// Calculate the third point
	cX := b[0] + r*math.Cos(angle)
	cY := b[1] + r*math.Sin(angle)

	return [2]float64{cX, cY}
}
