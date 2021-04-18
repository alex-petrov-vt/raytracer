package color_test

import (
	"testing"

	"github.com/alex-petrov-vt/raytracer/pkg/models/color"
	"github.com/alex-petrov-vt/raytracer/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestNewColor(t *testing.T) {
	c := color.NewColor(-0.5, 0.4, 1.7)
	assert.True(t, util.AlmostEqual(c.Red, -0.5))
	assert.True(t, util.AlmostEqual(c.Green, 0.4))
	assert.True(t, util.AlmostEqual(c.Blue, 1.7))
}

func TestEquals(t *testing.T) {
	c1 := color.NewColor(0.9, 0.6, 0.75)
	c2 := color.NewColor(0.9, 0.6, 0.75)
	c3 := color.NewColor(0.7, 0.1, 0.25)
	assert.True(t, color.Equals(c1, c2))
	assert.False(t, color.Equals(c1, c3))
	assert.False(t, color.Equals(c2, c3))
}

func TestAddingColors(t *testing.T) {
	c1 := color.NewColor(0.9, 0.6, 0.75)
	c2 := color.NewColor(0.7, 0.1, 0.25)
	c3 := color.Add(c1, c2)
	assert.True(t, color.Equals(c3, color.NewColor(1.6, 0.7, 1)))
}

func TestSubtractingColors(t *testing.T) {
	c1 := color.NewColor(0.9, 0.6, 0.75)
	c2 := color.NewColor(0.7, 0.1, 0.25)
	c3 := color.Subtract(c1, c2)
	assert.True(t, color.Equals(c3, color.NewColor(0.2, 0.5, 0.5)))
}

func TestMultiplingColorByScalar(t *testing.T) {
	c1 := color.NewColor(0.2, 0.3, 0.4)
	c1.Scale(2)
	assert.True(t, color.Equals(c1, color.NewColor(0.4, 0.6, 0.8)))
}

func TestMultiplyingTwoColors(t *testing.T) {
	c1 := color.NewColor(1, 0.2, 0.4)
	c2 := color.NewColor(0.9, 1, 0.1)
	c3 := color.Multiply(c1, c2)
	assert.True(t, color.Equals(c3, color.NewColor(0.9, 0.2, 0.04)))
}
