package vector

import (
	"errors"
	"math"

	"github.com/alex-petrov-vt/raytracer/pkg/util"
)

// Vector represents a set of three coordinates and the forth coordinate W which
// signifies if Vector is a Point (1) or a Vector(0)
type Vector struct {
	X, Y, Z, W float64
}

// NewPoint creates a point Vector from 3D coordinates
func NewPoint(x, y, z float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z,
		W: 1.0,
	}
}

// NewVector creates a vector Vector from 3D coordinates
func NewVector(x, y, z float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z,
		W: 0.0,
	}
}

// NewRawVector creates a vector from 4D coordinates
func New4DVector(x, y, z, w float64) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z,
		W: w,
	}
}

// AsSlice returns vector as a slice
func AsSlice(v *Vector) []float64 {
	return []float64{v.X, v.Y, v.Z, v.W}
}

// FromSlice creates a 4D vector from slice
func FromSlice(s []float64) (*Vector, error) {
	if len(s) != 4 {
		return nil, errors.New("can't create a 4D vector from a slice of length not equal to 4")
	}
	return &Vector{
		s[0],
		s[1],
		s[2],
		s[3],
	}, nil
}

// IsPoint returns true if Vector represents a Point
func IsPoint(v *Vector) bool {
	return util.FloatEquals(v.W, 1.0)
}

// IsVector returns true if Vector represents a vector
func IsVector(v *Vector) bool {
	return util.FloatEquals(v.W, 0.0)
}

// Equals compares two Vectors for equality
func Equals(v1, v2 *Vector) bool {
	return util.FloatEquals(v1.X, v2.X) && util.FloatEquals(v1.Y, v2.Y) &&
		util.FloatEquals(v1.Z, v2.Z) && util.FloatEquals(v1.W, v2.W)
}

// Add adds two Vectors together and returns the resulting Vector
func Add(t1, t2 *Vector) *Vector {
	return &Vector{
		t1.X + t2.X,
		t1.Y + t2.Y,
		t1.Z + t2.Z,
		t1.W + t2.W,
	}
}

// Subtract subtracts second Vector from the first Vector and returns the resulting
// Vector (or Point)
func Subtract(v1, v2 *Vector) *Vector {
	return &Vector{
		v1.X - v2.X,
		v1.Y - v2.Y,
		v1.Z - v2.Z,
		v1.W - v2.W,
	}
}

// Negate inverts all of the signs of the Vector coordinates
func Negate(v *Vector) *Vector {
	return &Vector{
		-v.X,
		-v.Y,
		-v.Z,
		-v.W,
	}
}

// Multiply multiples Vector by a scalar
func Multiply(v *Vector, s float64) *Vector {
	return &Vector{
		v.X * s,
		v.Y * s,
		v.Z * s,
		v.W * s,
	}
}

// Divide divides Vector by a scalar
func Divide(v *Vector, s float64) (*Vector, error) {
	if util.FloatEquals(s, 0) {
		return nil, errors.New("attempt to divide vector by 0")
	}
	return &Vector{
		v.X / s,
		v.Y / s,
		v.Z / s,
		v.W / s,
	}, nil
}

// Magnitude returns the magnitude of the vector (sqrt of squares of its terms)
func Magnitude(v *Vector) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)
}

// Normalize scales vector t to be unit vector in its direction
func Normalize(v *Vector) (*Vector, error) {
	m := Magnitude(v)
	if util.FloatEquals(m, 0) {
		return nil, errors.New("attempt to normalize zero vector")
	}
	return &Vector{
		v.X / m,
		v.Y / m,
		v.Z / m,
		v.W / m,
	}, nil
}

// Dot returns a dot product for vectors t1 and t2
func Dot(v1, v2 *Vector) float64 {
	return v1.X*v2.X + v1.Y*v2.Y + v1.Z*v2.Z + v1.W*v2.W
}

// Cross returns a cross product vector for vectors t1 and t2
func Cross(v1, v2 *Vector) *Vector {
	return NewVector(v1.Y*v2.Z-v1.Z*v2.Y, v1.Z*v2.X-v1.X*v2.Z, v1.X*v2.Y-v1.Y*v2.X)
}
