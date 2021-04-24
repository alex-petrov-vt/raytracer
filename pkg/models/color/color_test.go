package color

import (
	"testing"

	"github.com/alex-petrov-vt/raytracer/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestNewColor(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)
	assert.True(t, util.FloatEquals(c.Red, -0.5))
	assert.True(t, util.FloatEquals(c.Green, 0.4))
	assert.True(t, util.FloatEquals(c.Blue, 1.7))
}

func TestEquals(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.9, 0.6, 0.75)
	c3 := NewColor(0.7, 0.1, 0.25)
	assert.True(t, Equals(c1, c2))
	assert.False(t, Equals(c1, c3))
	assert.False(t, Equals(c2, c3))
}

func TestAddingColors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	c3 := Add(c1, c2)
	assert.True(t, Equals(c3, NewColor(1.6, 0.7, 1)))
}

func TestSubtractingColors(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	c3 := Subtract(c1, c2)
	assert.True(t, Equals(c3, NewColor(0.2, 0.5, 0.5)))
}

func TestMultiplingColorByScalar(t *testing.T) {
	c1 := NewColor(0.2, 0.3, 0.4)
	assert.True(t, Equals(Scale(c1, 2), NewColor(0.4, 0.6, 0.8)))
}

func TestMultiplyingTwoColors(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)
	c3 := Multiply(c1, c2)
	assert.True(t, Equals(c3, NewColor(0.9, 0.2, 0.04)))
}

func TestColorTo255Range(t *testing.T) {
	tests := map[string]struct {
		input *Color
		want  *Color
	}{
		"simple":    {input: NewColor(0.5, 0, 0), want: NewColor(128, 0, 0)},
		"overflow":  {input: NewColor(1.5, 0, 0), want: NewColor(255, 0, 0)},
		"underflow": {input: NewColor(-0.5, 0, 1), want: NewColor(0, 0, 255)},
	}

	for name, tc := range tests {
		got := ColorTo255Range(tc.input)
		if !Equals(got, tc.want) {
			t.Fatalf("%s: expected: %v, got %v", name, tc.want, got)
		}
	}
}
