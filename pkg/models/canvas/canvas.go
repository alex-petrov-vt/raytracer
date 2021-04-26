package canvas

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode/utf8"

	"github.com/alex-petrov-vt/raytracer/pkg/models/color"
)

const (
	ppmMagicNumber = "P3"
	colorRange     = 255
	maxLineSize    = 70
)

// Canvas is a representation of a screen of dimensions widht x height where all
// of pixels' colors are stored in internal 2D Colors slice.
type Canvas struct {
	Height, Width int
	Colors        [][]*color.Color
}

// NewCanvas creates new canvas of dimensions w x h (widht x height) and initializes
// all pixels to black color.
func NewCanvas(w, h int) *Canvas {
	colors := newColorMap(w, h)
	return &Canvas{
		Height: h,
		Width:  w,
		Colors: colors,
	}
}

func newColorMap(w, h int) [][]*color.Color {
	colors := make([][]*color.Color, h)
	for row := range colors {
		colors[row] = make([]*color.Color, w)
		for rowItem := range colors[row] {
			colors[row][rowItem] = color.NewColor(0, 0, 0)
		}
	}
	return colors
}

// WritePixel writes a color provided in col to the pixel located at [w][h]
func (c *Canvas) WritePixel(w, h int, col *color.Color) {
	c.Colors[h][w] = col
}

//  GetPixel returns a color of the pixel localted at [w][h]
func (c *Canvas) GetPixel(w, h int) *color.Color {
	return c.Colors[h][w]
}

// SaveToPPM saves canvas to a .ppm file
func (c *Canvas) SaveToPPM(file string) error {
	handle, err := os.Open(file)
	if err != nil {
		return err
	}
	defer handle.Close()
	return writePPM(handle, c)
}

func writePPM(handle io.Writer, c *Canvas) error {
	w := bufio.NewWriter(handle)
	if err := writeHeader(w, c); err != nil {
		return err
	}
	return writeData(w, c)
}

func writeHeader(w *bufio.Writer, c *Canvas) error {
	header := fmt.Sprintf("%s\n%d %d\n%d\n", ppmMagicNumber, c.Width, c.Height, colorRange)
	_, err := w.Write([]byte(header))
	if err != nil {
		return err
	}
	return w.Flush()
}

func writeData(w *bufio.Writer, c *Canvas) error {
	for _, row := range c.Colors {
		rowString := ""
		lineLength := 0
		for colorCount, currColor := range row {
			colorToSave := color.ColorTo255Range(currColor)
			colorToSaveString := fmt.Sprintf("%d %d %d", int(colorToSave.Red), int(colorToSave.Green), int(colorToSave.Blue))
			rowString += colorToSaveString
			lineLength += utf8.RuneCountInString(colorToSaveString)
			if lineLength > maxLineSize {
				rowString, lineLength = splitLongLine(rowString, lineLength)
			}
			if colorCount != len(row)-1 {
				rowString += " "
				lineLength += 1
			} else {
				rowString += "\n"
			}
		}

		if _, err := w.Write([]byte(rowString)); err != nil {
			return err
		}
	}
	return w.Flush()
}

func splitLongLine(line string, lineLength int) (string, int) {
	for i := utf8.RuneCountInString(line) - (lineLength - maxLineSize); i >= 0; i-- {
		if line[i] == ' ' {
			line = line[:i] + "\n" + line[i+1:]
			lineLength = len(line[i+1:])
			break
		}
	}
	return line, lineLength
}
