package triangle

import (
	"math"

	"github.com/Anaxarchus/zero-gdscript/pkg/vector3"
)

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

// Function to triangulate a point in 3D space
func TriangulateAngle3D(p1, p2 vector3.Vector3, angle float64) [3]float64 {
	// Calculate the direction vector from p1 to p2
	d := vector3.New(p2.X-p1.X, p2.Y-p1.Y, p2.Z-p1.Z)
	// Normalize the direction vector
	u := d.Normalized()
	// Calculate the distance between p1 and p2
	distance := p1.DistanceTo(p2)

	// Rotate the vector by the given angle in the plane defined by the direction vector
	// We use a simple rotation matrix in 2D for this purpose
	rotatedX := distance * math.Cos(angle)
	rotatedY := distance * math.Sin(angle)

	// Calculate the third point based on the rotation
	cx := p2.X + rotatedX*u.X
	cy := p2.Y + rotatedY*u.Y
	cz := p2.Z + rotatedY*u.Z

	return [3]float64{cx, cy, cz}
}
