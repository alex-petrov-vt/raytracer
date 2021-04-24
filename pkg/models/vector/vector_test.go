package vector

import (
	"errors"
	"math"
	"testing"

	"github.com/alex-petrov-vt/raytracer/pkg/util"

	"github.com/stretchr/testify/assert"
)

func TestNewPoint(t *testing.T) {
	vec := NewPoint(4.3, -4.2, 3.1)

	assert.True(t, util.FloatEquals(vec.X, 4.3))
	assert.True(t, util.FloatEquals(vec.Y, -4.2))
	assert.True(t, util.FloatEquals(vec.Z, 3.1))
	assert.True(t, util.FloatEquals(vec.W, 1.0))
	assert.True(t, IsPoint(vec))
	assert.False(t, IsVector(vec))
}

func TestNewVector(t *testing.T) {
	vec := NewVector(4.3, -4.2, 3.1)

	assert.True(t, util.FloatEquals(vec.X, 4.3))
	assert.True(t, util.FloatEquals(vec.Y, -4.2))
	assert.True(t, util.FloatEquals(vec.Z, 3.1))
	assert.True(t, util.FloatEquals(vec.W, 0.0))
	assert.False(t, IsPoint(vec))
	assert.True(t, IsVector(vec))
}

func TestVectorEquals(v *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(1, 2, 3)
	v3 := NewVector(1.1, 2, -3)
	assert.True(v, Equals(v1, v2))
	assert.False(v, Equals(v1, v3))
}

func TestAdd(t *testing.T) {
	v1 := NewPoint(3, -2, 5)
	v2 := NewVector(-2, 3, 1)
	t3 := Add(v1, v2)
	assert.True(t, Equals(t3, NewPoint(1, 1, 6)))

}

func TestSubtractPoints(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	p2 := NewPoint(5, 6, 7)
	p3 := Subtract(p1, p2)
	assert.True(t, Equals(p3, NewVector(-2, -4, -6)))
	assert.True(t, IsVector(p3))
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	v1 := NewVector(5, 6, 7)
	p2 := Subtract(p1, v1)
	assert.True(t, Equals(p2, NewPoint(-2, -4, -6)))
	assert.True(t, IsPoint(p2))
}

func TestSubtractVectors(t *testing.T) {
	v1 := NewVector(3, 2, 1)
	v2 := NewVector(5, 6, 7)
	v3 := Subtract(v1, v2)
	assert.True(t, Equals(v3, NewVector(-2, -4, -6)))
	assert.True(t, IsVector(v3))
}

func TestVectorNegation(t *testing.T) {
	v1 := NewVector(1, -2, 3)
	v2 := Negate(v1)
	assert.True(t, Equals(v2, NewVector(-1, 2, -3)))
}

func TestVectorMultiplicationByScalar(t *testing.T) {
	v1 := New4DVector(1, -2, 3, -4)
	v2 := Multiply(v1, 3.5)
	assert.True(t, Equals(v2, New4DVector(3.5, -7, 10.5, -14)))

	v3 := New4DVector(1, -2, 3, -4)
	v4 := Multiply(v3, 0.5)
	assert.True(t, Equals(v4, New4DVector(0.5, -1, 1.5, -2)))
}

func TestVectorDivisionByScalar(t *testing.T) {
	v1 := New4DVector(1, -2, 3, -4)
	v2, err := Divide(v1, 2)
	assert.True(t, Equals(v2, New4DVector(0.5, -1, 1.5, -2)))
	assert.Nil(t, err)
	_, err = Divide(v1, 0)
	assert.NotNil(t, err)
}

func TestVectorMagnitude(t *testing.T) {
	tests := map[string]struct {
		input *Vector
		want  float64
	}{
		"simple":                  {input: NewVector(1, 0, 0), want: 1},
		"all zero components":     {input: NewVector(0, 0, 0), want: 0},
		"all non-zero components": {input: NewVector(1, 2, 3), want: math.Sqrt(14)},
		"negative components":     {input: NewVector(-1, -2, -3), want: math.Sqrt(14)},
	}

	for name, tc := range tests {
		got := Magnitude(tc.input)
		if !assert.True(t, util.FloatEquals(tc.want, got)) {
			t.Fatalf("%s: expected: %v, got %v", name, tc.want, got)
		}
	}
}

func TestVectorNormalization(t *testing.T) {
	tests := map[string]struct {
		name  string
		input *Vector
		want  *Vector
		err   error
	}{
		"simple":                  {input: NewVector(4, 0, 0), want: NewVector(1, 0, 0), err: nil},
		"all non-zero components": {input: NewVector(1, 2, 3), want: NewVector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14)), err: nil},
		"all zero components":     {input: NewVector(0, 0, 0), want: nil, err: errors.New("")},
	}

	for name, tc := range tests {
		got, err := Normalize(tc.input)
		if err != nil {
			if tc.err == nil {
				t.Fatalf("%s: expected no error for input: %v, but got %v", name, tc.input, err)
			}
			continue
		}

		if err == nil && tc.err != nil {
			t.Fatalf("%s: expected error for input: %v, but got %v", name, tc.input, got)
		}

		if !Equals(tc.want, got) {
			t.Fatalf("%s: expected: %v, got %v", name, tc.want, got)
		}
	}
}

func TestDotProduct(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)
	assert.True(t, util.FloatEquals(Dot(v1, v2), 20))
}

func TestCrossProduct(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)
	assert.True(t, Equals(Cross(v1, v2), NewVector(-1, 2, -1)))
	assert.True(t, Equals(Cross(v2, v1), NewVector(1, -2, 1)))
}
