package canvas

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/alex-petrov-vt/raytracer/pkg/models/color"
	"github.com/stretchr/testify/assert"
)

func TestNewCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	assert.Equal(t, c.Width, 10)
	assert.Equal(t, c.Height, 20)

	for _, col := range c.Colors {
		for _, pix := range col {
			assert.True(t, color.Equals(pix, color.NewColor(0, 0, 0)))
		}
	}
}

func TestWritePixel(t *testing.T) {
	c := NewCanvas(10, 20)
	red := color.NewColor(1, 0, 0)

	c.WritePixel(2, 3, red)
	newRed := c.GetPixel(2, 3)
	assert.True(t, color.Equals(red, newRed))
}

func TestCanvasToPPM(t *testing.T) {
	colors := []*color.Color{color.NewColor(1.5, 0, 0), color.NewColor(0, 0.5, 0),
		color.NewColor(-0.5, 0, 1)}
	locations := []int{0, 0, 2, 1, 4, 2}
	c1 := prepareCanvas(5, 3, colors, locations)

	colors = nil
	// Create canvas with every pixel set to color (1, 0.8, 0.6) of size 10x2
	for i := 0; i < 10; i++ {
		for j := 0; j < 2; j++ {
			colors = append(colors, color.NewColor(1, 0.8, 0.6))
		}
	}
	locations = nil
	for j := 0; j < 2; j++ {
		for i := 0; i < 10; i++ {
			locations = append(locations, i)
			locations = append(locations, j)
		}
	}
	c2 := prepareCanvas(10, 2, colors, locations)

	tests := map[string]struct {
		input *Canvas
		want  string
	}{
		"simple": {input: c1, want: `P3
5 3
255
255 0 0 0 0 0 0 0 0 0 0 0 0 0 0
0 0 0 0 0 0 0 128 0 0 0 0 0 0 0
0 0 0 0 0 0 0 0 0 0 0 0 0 0 255
`},
		"long lines": {input: c2, want: `P3
10 2
255
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204
153 255 204 153 255 204 153 255 204 153 255 204 153
`},
	}

	for name, tc := range tests {
		var b bytes.Buffer
		writePPM(&b, tc.input)
		r := bufio.NewReader(&b)
		got := b.String()
		if tc.want != string(got) {
			t.Fatalf("%s: expected: %v, got %v", name, tc.want, got)
		}
		assert.Equal(t, r.Buffered(), 0)
	}
}

func prepareCanvas(w, h int, colors []*color.Color, locations []int) *Canvas {
	c := NewCanvas(w, h)
	i := 0
	for _, color := range colors {
		c.WritePixel(locations[i], locations[i+1], color)
		i += 2
	}
	return c
}
