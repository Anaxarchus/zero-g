package zerotriangles

import (
	"math"

	"github.com/Anaxarchus/zero-gdscript/pkg/vector3"
)

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
