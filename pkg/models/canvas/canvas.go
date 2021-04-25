package canvas

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode/utf8"

	"github.com/alex-petrov-vt/raytracer/pkg/models/color"
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
	colors := make([][]*color.Color, h)
	for j := range colors {
		colors[j] = make([]*color.Color, w)
		for i := range colors[j] {
			colors[j][i] = color.NewColor(0, 0, 0)
		}
	}

	return &Canvas{
		Height: h,
		Width:  w,
		Colors: colors,
	}
}

// WritePixel writes a color provided in col to the pixel located at [w][h]
func (c *Canvas) WritePixel(w, h int, col *color.Color) {
	c.Colors[h][w] = col
}

//  GetPixel returns a color of the pixel localted at [w][h]
func (c *Canvas) GetPixel(w, h int) *color.Color {
	return c.Colors[h][w]
}

func (c *Canvas) SaveToPPM(file string) error {
	handle, err := os.Open(file)
	if err != nil {
		return err
	}
	defer handle.Close()

	err = writePPM(handle, c)
	return err
}

func writePPM(handle io.Writer, c *Canvas) error {
	w := bufio.NewWriter(handle)

	err := writeHeader(w, c)
	if err != nil {
		return err
	}

	err = writeData(w, c)
	return err
}

func writeHeader(w *bufio.Writer, c *Canvas) error {
	header := fmt.Sprintf("P3\n%d %d\n255\n", c.Width, c.Height)
	_, err := w.Write([]byte(header))
	if err != nil {
		return err
	}
	err = w.Flush()
	return err
}

func writeData(w *bufio.Writer, c *Canvas) error {
	for _, row := range c.Colors {
		rowString := ""
		lineLenght := 0
		for rowCount, c := range row {
			colorToSave := color.ColorTo255Range(c)
			colorToSaveString := fmt.Sprintf("%d %d %d", int(colorToSave.Red), int(colorToSave.Green), int(colorToSave.Blue))
			rowString += colorToSaveString
			lineLenght += utf8.RuneCountInString(colorToSaveString)
			if lineLenght > 70 {
				for i := utf8.RuneCountInString(rowString) - (lineLenght - 70); i >= 0; i-- {
					if rowString[i] == ' ' {
						rowString = rowString[:i] + "\n" + rowString[i+1:]
						lineLenght = len(rowString[i+1:])
						break
					}
				}
			}
			if rowCount != len(row)-1 {
				rowString += " "
				lineLenght += 1
			} else {
				rowString += "\n"
			}
		}
		_, err := w.Write([]byte(rowString))
		if err != nil {
			return err
		}
	}
	err := w.Flush()
	return err
}
