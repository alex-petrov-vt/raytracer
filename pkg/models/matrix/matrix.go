package matrix

import (
	"errors"

	"github.com/alex-petrov-vt/raytracer/pkg/models/vector"
	"github.com/alex-petrov-vt/raytracer/pkg/util"
)

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
			if !util.FloatEquals(m1.Elements[row][col], m2.Elements[row][col]) {
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

// MultiplyByVector multiplies a matrix by a vector
func MultiplyByVector(m *Matrix, v *vector.Vector) (*vector.Vector, error) {
	if m.Width != 4 {
		return nil, errors.New("incompatible dimensions for matrix multiplication")
	}

	var result []float64
	vAsSlice := vector.AsSlice(v)
	for row := range m.Elements {
		rowResult := 0.0
		for i := 0; i < 4; i++ {
			rowResult += m.Elements[row][i] * vAsSlice[i]
		}
		result = append(result, rowResult)
	}
	return vector.FromSlice(result)
}

// Transpose transposes the matrix (turns rows into cols)
func Transpose(m *Matrix) *Matrix {
	var result [][]float64
	for col := 0; col < m.Height; col++ {
		var newRow []float64
		for row := 0; row < m.Width; row++ {
			newRow = append(newRow, m.Elements[row][col])
		}
		result = append(result, newRow)
	}
	return NewMatrix(result)
}

// GetDeterminant returns a determinant of the matrix
func GetDeterminant(m *Matrix) float64 {
	result := 0.0
	if m.Width == 2 && m.Height == 2 {
		a, _ := m.GetElement(0, 0)
		d, _ := m.GetElement(1, 1)
		b, _ := m.GetElement(1, 0)
		c, _ := m.GetElement(0, 1)
		result = a*d - b*c
	} else {
		for col := 0; col < m.Width; col++ {
			result += m.Elements[0][col] * GetCofactor(m, 0, col)
		}
	}
	return result
}

// GetSubmatrix returns a submatrix of matrix m with skipRow and skipCol removed
func GetSubmatrix(m *Matrix, skipRow, skipCol int) *Matrix {
	var result [][]float64
	for row := 0; row < m.Height; row++ {
		if row == skipRow {
			continue
		}
		var newRow []float64
		for col := 0; col < m.Width; col++ {
			if col == skipCol {
				continue
			}
			newRow = append(newRow, m.Elements[row][col])
		}
		result = append(result, newRow)
	}
	return NewMatrix(result)
}

// GetMinor returns a minor of a matrix element at (row, col)
func GetMinor(m *Matrix, row, col int) float64 {
	sm := GetSubmatrix(m, row, col)
	return GetDeterminant(sm)
}

func GetCofactor(m *Matrix, row, col int) float64 {
	minor := GetMinor(m, row, col)
	if (row+col)%2 == 0 {
		return minor
	} else {
		return -minor
	}
}

func IsInvertible(m *Matrix) bool {
	return GetDeterminant(m) != 0
}

func GetInverse(m *Matrix) (*Matrix, error) {
	if !IsInvertible(m) {
		return nil, errors.New("matrix is not invertible")
	}
	var result [][]float64
	det := GetDeterminant(m)
	for row := 0; row < m.Height; row++ {
		var newRow []float64
		for col := 0; col < m.Width; col++ {
			c := GetCofactor(m, row, col)
			newRow = append(newRow, c/det)
		}
		result = append(result, newRow)
	}
	return Transpose(NewMatrix(result)), nil
}
