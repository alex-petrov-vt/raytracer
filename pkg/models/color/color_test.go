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

func TestAddingColors(t *testing.T) {
	c1 := color.NewColor(0.9, 0.6, 0.75)
	c2 := color.NewColor(0.7, 0.1, 0.25)
	c3 := color.Add(c1, c2)

	assert.True(t, util.AlmostEqual(c3.Red, 1.6))
	assert.True(t, util.AlmostEqual(c3.Green, 0.7))
	assert.True(t, util.AlmostEqual(c3.Blue, 1))
}

func TestSubtractingColors(t *testing.T) {
	c1 := color.NewColor(0.9, 0.6, 0.75)
	c2 := color.NewColor(0.7, 0.1, 0.25)
	c3 := color.Subtract(c1, c2)

	assert.True(t, util.AlmostEqual(c3.Red, 0.2))
	assert.True(t, util.AlmostEqual(c3.Green, 0.5))
	assert.True(t, util.AlmostEqual(c3.Blue, 0.5))
}

func TestMultiplingColorByScalar(t *testing.T) {
	c1 := color.NewColor(0.2, 0.3, 0.4)
	c1.Scale(2)
	assert.True(t, util.AlmostEqual(c1.Red, 0.4))
	assert.True(t, util.AlmostEqual(c1.Green, 0.6))
	assert.True(t, util.AlmostEqual(c1.Blue, 0.8))
}

func TestMultiplyingTwoColors(t *testing.T) {
	c1 := color.NewColor(1, 0.2, 0.4)
	c2 := color.NewColor(0.9, 1, 0.1)
	c3 := color.Multiply(c1, c2)
	assert.True(t, util.AlmostEqual(c3.Red, 0.9))
	assert.True(t, util.AlmostEqual(c3.Green, 0.2))
	assert.True(t, util.AlmostEqual(c3.Blue, 0.04))
}
