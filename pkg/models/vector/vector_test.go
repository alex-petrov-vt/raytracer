package vector_test

import (
	"math"
	"testing"

	"github.com/alex-petrov-vt/raytracer/pkg/models/vector"
	"github.com/alex-petrov-vt/raytracer/pkg/util"

	"github.com/stretchr/testify/assert"
)

func TestPointTuple(t *testing.T) {
	tup := vector.NewTuple(4.3, -4.2, 3.1, 1.0)

	assert.True(t, util.AlmostEqual(tup.X, 4.3))
	assert.True(t, util.AlmostEqual(tup.Y, -4.2))
	assert.True(t, util.AlmostEqual(tup.Z, 3.1))
	assert.True(t, util.AlmostEqual(tup.W, 1.0))
	assert.True(t, tup.IsPoint())
	assert.False(t, tup.IsVector())
}

func TestVectorTuple(t *testing.T) {
	tup := vector.NewTuple(4.3, -4.2, 3.1, 0.0)

	assert.True(t, util.AlmostEqual(tup.X, 4.3))
	assert.True(t, util.AlmostEqual(tup.Y, -4.2))
	assert.True(t, util.AlmostEqual(tup.Z, 3.1))
	assert.True(t, util.AlmostEqual(tup.W, 0.0))
	assert.False(t, tup.IsPoint())
	assert.True(t, tup.IsVector())
}

func TestNewPoint(t *testing.T) {
	p := vector.NewPoint(4, -4, 3)

	assert.True(t, util.AlmostEqual(p.X, 4))
	assert.True(t, util.AlmostEqual(p.Y, -4))
	assert.True(t, util.AlmostEqual(p.Z, 3))
	assert.True(t, util.AlmostEqual(p.W, 1.0))
	assert.True(t, p.IsPoint())
	assert.False(t, p.IsVector())
}

func TestNewVector(t *testing.T) {
	p := vector.NewVector(4, -4, 3)

	assert.True(t, util.AlmostEqual(p.X, 4))
	assert.True(t, util.AlmostEqual(p.Y, -4))
	assert.True(t, util.AlmostEqual(p.Z, 3))
	assert.True(t, util.AlmostEqual(p.W, 0.0))
	assert.False(t, p.IsPoint())
	assert.True(t, p.IsVector())
}

func TestEquals(t *testing.T) {
	t1 := vector.NewTuple(1, 2, 3, 0)
	t2 := vector.NewTuple(1, 2, 3, 0)
	t3 := vector.NewTuple(1, 2, 3, 1)
	t4 := vector.NewTuple(1.1, 2, 3, 0)
	assert.True(t, vector.Equals(t1, t2))
	assert.False(t, vector.Equals(t1, t3))
	assert.False(t, vector.Equals(t1, t4))
}

func TestAdd(t *testing.T) {
	t1 := vector.NewTuple(3, -2, 5, 1)
	t2 := vector.NewTuple(-2, 3, 1, 0)
	t3 := vector.Add(t1, t2)
	assert.True(t, vector.Equals(t3, vector.NewTuple(1, 1, 6, 1)))

}

func TestSubtractPoints(t *testing.T) {
	p1 := vector.NewPoint(3, 2, 1)
	p2 := vector.NewPoint(5, 6, 7)
	p3 := vector.Subtract(p1, p2)
	assert.True(t, vector.Equals(p3, vector.NewTuple(-2, -4, -6, 0)))
	assert.True(t, p3.IsVector())
}

func TestSubtractVectorFromPoint(t *testing.T) {
	p1 := vector.NewPoint(3, 2, 1)
	v1 := vector.NewVector(5, 6, 7)
	p2 := vector.Subtract(p1, v1)
	assert.True(t, vector.Equals(p2, vector.NewTuple(-2, -4, -6, 1)))
	assert.True(t, p2.IsPoint())
}

func TestSubtractVectors(t *testing.T) {
	v1 := vector.NewVector(3, 2, 1)
	v2 := vector.NewVector(5, 6, 7)
	v3 := vector.Subtract(v1, v2)
	assert.True(t, vector.Equals(v3, vector.NewTuple(-2, -4, -6, 0)))
	assert.True(t, v3.IsVector())
}

func TestTupleNegation(t *testing.T) {
	t1 := vector.NewTuple(1, -2, 3, -4)
	t1.Negate()
	assert.True(t, vector.Equals(t1, vector.NewTuple(-1, 2, -3, 4)))
}

func TestTupleMultiplicationByScalar(t *testing.T) {
	t1 := vector.NewTuple(1, -2, 3, -4)
	t1.Multiply(3.5)
	assert.True(t, vector.Equals(t1, vector.NewTuple(3.5, -7, 10.5, -14)))

	t2 := vector.NewTuple(1, -2, 3, -4)
	t2.Multiply(0.5)
	assert.True(t, vector.Equals(t2, vector.NewTuple(0.5, -1, 1.5, -2)))
}

func TestTupleDivisionByScalar(t *testing.T) {
	t1 := vector.NewTuple(1, -2, 3, -4)
	t1.Divide(2)
	assert.True(t, vector.Equals(t1, vector.NewTuple(0.5, -1, 1.5, -2)))
}

func TestVectorMagnitude(t *testing.T) {
	v1 := vector.NewVector(1, 0, 0)
	assert.True(t, util.AlmostEqual(v1.Magnitude(), 1))
	v2 := vector.NewVector(0, 1, 0)
	assert.True(t, util.AlmostEqual(v2.Magnitude(), 1))
	v3 := vector.NewVector(0, 0, 1)
	assert.True(t, util.AlmostEqual(v3.Magnitude(), 1))
	v4 := vector.NewVector(1, 2, 3)
	assert.True(t, util.AlmostEqual(v4.Magnitude(), math.Sqrt(14)))
	v5 := vector.NewVector(-1, -2, -3)
	assert.True(t, util.AlmostEqual(v5.Magnitude(), math.Sqrt(14)))
}

func TestVectorNormalization(t *testing.T) {
	v1 := vector.NewVector(4, 0, 0)
	v1.Normalize()
	assert.True(t, vector.Equals(v1, vector.NewVector(1, 0, 0)))
	v2 := vector.NewVector(1, 2, 3)
	v2.Normalize()
	assert.True(t, vector.Equals(v2, vector.NewVector(1/math.Sqrt(14), 2/math.Sqrt(14), 3/math.Sqrt(14))))

	util.AlmostEqual(v1.Magnitude(), 1)
	util.AlmostEqual(v2.Magnitude(), 1)
}

func TestDotProduct(t *testing.T) {
	v1 := vector.NewVector(1, 2, 3)
	v2 := vector.NewVector(2, 3, 4)
	assert.True(t, util.AlmostEqual(vector.Dot(v1, v2), 20))
}

func TestCrossProduct(t *testing.T) {
	v1 := vector.NewVector(1, 2, 3)
	v2 := vector.NewVector(2, 3, 4)
	assert.True(t, vector.Equals(vector.Cross(v1, v2), vector.NewVector(-1, 2, -1)))
	assert.True(t, vector.Equals(vector.Cross(v2, v1), vector.NewVector(1, -2, 1)))
}
