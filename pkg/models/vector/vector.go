package vector

import (
	"math"

	"github.com/alex-petrov-vt/raytracer/pkg/util"
)

// Tuple represents a set of three coordinates and the forth coordinate W which
// signifies if Tuple is a Point (1) or a Vector(0)
type Tuple struct {
	X, Y, Z, W float64
}

// NewTuple creates new tuple from coordinates and W value
func NewTuple(x, y, z, w float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

// NewPoint creates a point tuple from 3D coordinates
func NewPoint(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 1.0,
	}
}

// NewVector creates a vector tuple from 3D coordinates
func NewVector(x, y, z float64) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
		Z: z,
		W: 0.0,
	}
}

// IsPoint returns true if tuple represents a Point
func (t *Tuple) IsPoint() bool {
	return util.AlmostEqual(t.W, 1.0)
}

// IsVector returns true if Tuple represents a vector
func (t *Tuple) IsVector() bool {
	return util.AlmostEqual(t.W, 0.0)
}

// Equals compares two tuples for equality
func Equals(t1, t2 *Tuple) bool {
	return t1.X == t2.X && t1.Y == t2.Y && t1.Z == t2.Z && t1.W == t2.W
}

// Add adds two tuples together and returns the resulting tuple
func Add(t1, t2 *Tuple) *Tuple {
	return &Tuple{
		t1.X + t2.X,
		t1.Y + t2.Y,
		t1.Z + t2.Z,
		t1.W + t2.W,
	}
}

// Subtract subtracts second tuple from the first tuple and returns the resulting
// tuple
func Subtract(t1, t2 *Tuple) *Tuple {
	return &Tuple{
		t1.X - t2.X,
		t1.Y - t2.Y,
		t1.Z - t2.Z,
		t1.W - t2.W,
	}
}

// Negate inverts all of the signs of the tuple coordinates
func (t *Tuple) Negate() {
	t.X = -t.X
	t.Y = -t.Y
	t.Z = -t.Z
	t.W = -t.W
}

// Multiply multiples tuple by a scalar
func (t *Tuple) Multiply(s float64) {
	t.X = t.X * s
	t.Y = t.Y * s
	t.Z = t.Z * s
	t.W = t.W * s
}

// Multiply multiples tuple by a scalar
func (t *Tuple) Divide(s float64) {
	t.X = t.X / s
	t.Y = t.Y / s
	t.Z = t.Z / s
	t.W = t.W / s
}

// Magnitude returns the magnitude of the vector (sqrt of squares of its terms)
func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W)
}

// Normalize scales vector t to be unit vector in its direction
func (t *Tuple) Normalize() {
	m := t.Magnitude()
	t.X = t.X / m
	t.Y = t.Y / m
	t.Z = t.Z / m
}

// Dot returns a dot product for vectors t1 and t2
func Dot(t1, t2 *Tuple) float64 {
	return t1.X*t2.X + t1.Y*t2.Y + t1.Z*t2.Z + t1.W*t2.W
}

// Cross returns a cross product vector for vectors t1 and t2
func Cross(t1, t2 *Tuple) *Tuple {
	return NewVector(t1.Y*t2.Z-t1.Z*t2.Y, t1.Z*t2.X-t1.X*t2.Z, t1.X*t2.Y-t1.Y*t2.X)
}
