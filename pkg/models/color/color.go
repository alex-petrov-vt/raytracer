package color

import "github.com/alex-petrov-vt/raytracer/pkg/util"

// Color is a struct that represents a color in an RGB format
type Color struct {
	Red, Green, Blue float64
}

// NewColor creates new color from red, green, and blue values
func NewColor(r, g, b float64) *Color {
	return &Color{r, g, b}
}

// Add adds two colors together
func Add(c1, c2 *Color) *Color {
	return &Color{c1.Red + c2.Red, c1.Green + c2.Green, c1.Blue + c2.Blue}
}

// Subtract subtracts second color from the first
func Subtract(c1, c2 *Color) *Color {
	return &Color{c1.Red - c2.Red, c1.Green - c2.Green, c1.Blue - c2.Blue}
}

// Scale scales a color by a scalar
func Scale(c *Color, s float64) *Color {
	return &Color{
		c.Red * s,
		c.Green * s,
		c.Blue * s,
	}
}

// Multiply computes Hadamard (or Schur) product of two colors
func Multiply(c1, c2 *Color) *Color {
	return &Color{c1.Red * c2.Red, c1.Green * c2.Green, c1.Blue * c2.Blue}
}

func Equals(c1, c2 *Color) bool {
	return util.FloatEquals(c1.Red, c2.Red) && util.FloatEquals(c1.Green, c2.Green) &&
		util.FloatEquals(c1.Blue, c2.Blue)
}
