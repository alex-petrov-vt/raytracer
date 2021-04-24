package canvas

import (
	"bufio"
	"bytes"
	"fmt"
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
	c := NewCanvas(5, 3)
	c1 := color.NewColor(1.5, 0, 0)
	c2 := color.NewColor(0, 0.5, 0)
	c3 := color.NewColor(-0.5, 0, 1)
	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)
	var b bytes.Buffer
	writePPM(&b, c)
	r := bufio.NewReader(&b)

	// Header
	magicNum, err := r.ReadBytes('\n')
	assert.Nil(t, err)
	assert.Equal(t, "P3\n", string(magicNum))
	widthHeight, err := r.ReadBytes('\n')
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf("%d %d\n", c.Width, c.Height), string(widthHeight))
	maxColor, err := r.ReadBytes('\n')
	assert.Nil(t, err)
	assert.Equal(t, "255\n", string(maxColor))

	// Data
	row1, err := r.ReadBytes('\n')
	assert.Nil(t, err)
	assert.Equal(t, "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n", string(row1))
	row2, err := r.ReadBytes('\n')
	assert.Nil(t, err)
	assert.Equal(t, "0 0 0 0 0 0 0 128 0 0 0 0 0 0 0\n", string(row2))
	row3, err := r.ReadBytes('\n')
	assert.Nil(t, err)
	assert.Equal(t, "0 0 0 0 0 0 0 0 0 0 0 0 0 0 255\n", string(row3))
}
