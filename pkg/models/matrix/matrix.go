package matrix

import "errors"

type Matrix struct {
	Width, Height int
	Elements      [][]float64
}

// NewMatrix creates new matrix given an input 2D array
func NewMatrix(input [][]float64) *Matrix {
	if len(input) == 0 {
		return &Matrix{
			len(input),
			0,
			input,
		}
	}
	return &Matrix{
		len(input),
		len(input[0]),
		input,
	}
}

// GetElement returns an element with coordinates [row][col] or error if those
// are outside of the boundaries of the matrix
func (m *Matrix) GetElement(row, col int) (float64, error) {
	if row < 0 || row >= m.Height || col < 0 || col >= m.Width {
		return 0, errors.New("access elements outside of matrix")
	}
	return m.Elements[row][col], nil
}

// Equals compares two matrices for equality
func Equals(m1, m2 *Matrix) bool {
	if m1.Height != m2.Height || m1.Width != m2.Width {
		return false
	}

	for row := range m1.Elements {
		for col := range m1.Elements[row] {
			if m1.Elements[row][col] != m2.Elements[row][col] {
				return false
			}
		}
	}
	return true
}

//Multiply multiplies two matrices together
func Multiply(m1, m2 *Matrix) (*Matrix, error) {
	if m1.Height == 0 || m1.Width == 0 || m2.Height == 0 || m2.Width == 0 {
		return nil, errors.New("one or both of matrices are empty")
	}
	if m1.Width != m2.Height {
		return nil, errors.New("incompatible dimensions for matrix multiplication")
	}

	var newElems [][]float64
	for row := range m1.Elements {
		var newRow []float64
		for col := range m1.Elements[row] {
			result := 0.0
			for i := range m2.Elements {
				result += m1.Elements[row][i] * m2.Elements[i][col]
			}
			newRow = append(newRow, result)
		}
		newElems = append(newElems, newRow)
	}
	return NewMatrix(newElems), nil
}
