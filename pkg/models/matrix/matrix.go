package matrix

import "errors"

type Matrix struct {
	elements [][]float64
}

// NewMatrix creates new matrix given an input 2D array
func NewMatrix(input [][]float64) *Matrix {
	return &Matrix{
		input,
	}
}

// GetElement returns an element with coordinates [row][col] or error if those
// are outside of the boundaries of the matrix
func (m *Matrix) GetElement(row, col int) (float64, error) {
	if row < 0 || row >= len(m.elements) || col < 0 || col >= len(m.elements[0]) {
		return 0, errors.New("access elements outside of matrix")
	}
	return m.elements[row][col], nil
}

// Equals compares two matrices for equality
func Equals(m1, m2 *Matrix) bool {
	if len(m1.elements) != len(m2.elements) {
		return false
	}
	if len(m1.elements) > 0 && len(m2.elements) > 0 && len(m1.elements[0]) != len(m2.elements[0]) {
		return false
	}

	for row := range m1.elements {
		for col := range m1.elements[row] {
			if m1.elements[row][col] != m2.elements[row][col] {
				return false
			}
		}
	}
	return true
}

//Multiply multiplies two matrices together
func Multiply(m1, m2 *Matrix) (*Matrix, error) {
	if len(m1.elements) == 0 || len(m2.elements) == 0 || len(m1.elements[0]) == 0 || len(m2.elements[0]) == 0 {
		return nil, errors.New("one or both of matrices are empty")
	}
	if len(m1.elements) != len(m2.elements[0]) {
		return nil, errors.New("incompatible dimensions for matrix multiplication")
	}

	var newElems [][]float64
	for row := range m1.elements {
		var newRow []float64
		for col := range m1.elements[row] {
			result := 0.0
			for i := range m2.elements {
				result += m1.elements[row][i] * m2.elements[i][col]
			}
			newRow = append(newRow, result)
		}
		newElems = append(newElems, newRow)
	}
	return NewMatrix(newElems), nil
}
