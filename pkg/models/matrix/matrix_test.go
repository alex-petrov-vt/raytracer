package matrix

import (
	"errors"
	"testing"

	"github.com/alex-petrov-vt/raytracer/pkg/models/vector"
	"github.com/alex-petrov-vt/raytracer/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestMatrices(t *testing.T) {
	m1 := NewMatrix([][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	})
	m2 := NewMatrix([][]float64{
		{-3, 5},
		{1, -2},
	})
	m3 := NewMatrix([][]float64{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	})
	m4 := NewMatrix([][]float64{{}})

	tests := []struct {
		m    *Matrix
		row  int
		col  int
		want float64
		err  error
	}{
		{m: m1, row: 0, col: 0, want: 1.0, err: nil},
		{m: m1, row: 0, col: 3, want: 4, err: nil},
		{m: m1, row: 1, col: 0, want: 5.5, err: nil},
		{m: m1, row: 1, col: 2, want: 7.5, err: nil},
		{m: m1, row: 2, col: 2, want: 11, err: nil},
		{m: m1, row: 3, col: 0, want: 13.5, err: nil},
		{m: m1, row: 3, col: 2, want: 15.5, err: nil},
		{m: m1, row: 100, col: 100, want: 0, err: errors.New("")},
		{m: m2, row: 0, col: 0, want: -3, err: nil},
		{m: m2, row: 0, col: 1, want: 5, err: nil},
		{m: m2, row: 1, col: 0, want: 1, err: nil},
		{m: m2, row: 1, col: 1, want: -2, err: nil},
		{m: m2, row: -1, col: 1, want: 0, err: errors.New("")},
		{m: m3, row: 0, col: 0, want: -3, err: nil},
		{m: m3, row: 1, col: 1, want: -2, err: nil},
		{m: m3, row: 2, col: 2, want: 1, err: nil},
		{m: m4, row: 0, col: 0, want: 0, err: errors.New("")},
		{m: m4, row: -1, col: 0, want: 0, err: errors.New("")},
	}

	for _, tc := range tests {
		got, err := tc.m.GetElement(tc.row, tc.col)
		if err != nil {
			if tc.err == nil {
				t.Fatalf("unexpected error for input [%v][%v]: %v", tc.row, tc.col, err)
			}
			continue
		}

		if !util.FloatEquals(got, tc.want) {
			t.Fatalf("expected: %v, got %v", tc.want, got)
		}
	}
}

func TestMatrixEqual(t *testing.T) {
	a := NewMatrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	b := NewMatrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	c := NewMatrix([][]float64{
		{2, 3, 4, 5},
		{6, 7, 8, 9},
		{8, 7, 6, 5},
		{4, 3, 2, 1},
	})
	d := NewMatrix([][]float64{{}})
	e := NewMatrix([][]float64{{}})
	f := NewMatrix([][]float64{
		{1, 2},
		{3, 4},
	})
	g := NewMatrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	})
	h := NewMatrix([][]float64{
		{1},
		{2},
		{3},
		{4},
	})

	tests := map[string]struct {
		m1   *Matrix
		m2   *Matrix
		want bool
	}{
		"simple equal":          {m1: a, m2: b, want: true},
		"simple not equal":      {m1: a, m2: c, want: false},
		"empty equal":           {m1: d, m2: e, want: true},
		"empty not equal":       {m1: a, m2: d, want: false},
		"different dimensions":  {m1: a, m2: f, want: false},
		"different num of rows": {m1: a, m2: g, want: false},
		"different num of cols": {m1: a, m2: h, want: false},
	}

	for name, tc := range tests {
		got := IsEqual(tc.m1, tc.m2)
		if got != tc.want {
			t.Fatalf("%s: exptected %v, got %v", name, tc.want, got)
		}
	}
}

func TestMatrixMultiplication(t *testing.T) {
	a := NewMatrix([][]float64{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 8, 7, 6},
		{5, 4, 3, 2},
	})
	b := NewMatrix([][]float64{
		{-2, 1, 2, 3},
		{3, 2, 1, -1},
		{4, 3, 6, 5},
		{1, 2, 7, 8},
	})
	c := NewMatrix([][]float64{
		{20, 22, 50, 48},
		{44, 54, 114, 108},
		{40, 58, 110, 102},
		{16, 26, 46, 42},
	})
	d := NewMatrix([][]float64{{}})
	e := NewMatrix([][]float64{
		{1, 2},
		{3, 4},
	})
	f := NewMatrix([][]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})

	tests := map[string]struct {
		m1   *Matrix
		m2   *Matrix
		want *Matrix
		err  error
	}{
		"simple":                  {m1: a, m2: b, want: c, err: nil},
		"by identity matrix":      {m1: a, m2: f, want: a, err: nil},
		"one empty":               {m1: a, m2: d, want: nil, err: errors.New("")},
		"incompatible dimensions": {m1: a, m2: e, want: nil, err: errors.New("")},
	}

	for name, tc := range tests {
		got, err := Multiply(tc.m1, tc.m2)
		if err != nil {
			if tc.err == nil {
				t.Fatalf("%s: unexpected error %v", name, err)
			}
			continue
		}

		if !IsEqual(got, tc.want) {
			t.Fatalf("%s: expected %v, got %v", name, tc.want, got)
		}
	}
}

func TestMatrixMultiplicationByVector(t *testing.T) {
	a := NewMatrix([][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})
	b := NewMatrix([][]float64{
		{1, 2},
		{3, 4},
	})
	c := vector.New4DVector(1, 2, 3, 1)
	d := NewMatrix([][]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})

	tests := map[string]struct {
		m    *Matrix
		v    *vector.Vector
		want *vector.Vector
		err  error
	}{
		"simple":                  {m: a, v: c, want: vector.New4DVector(18, 24, 33, 1), err: nil},
		"by identity matrix":      {m: d, v: c, want: c, err: nil},
		"incompatible dimensions": {m: b, v: c, want: nil, err: errors.New("")},
	}

	for name, tc := range tests {
		got, err := MultiplyByVector(tc.m, tc.v)
		if err != nil {
			if tc.err == nil {
				t.Fatalf("%s: unexpected error %v", name, err)
			}
			continue
		}

		if !vector.Equals(got, tc.want) {
			t.Fatalf("%s: expected %v, got %v", name, tc.want, got)
		}
	}
}

func TestMatrixTransposition(t *testing.T) {
	a := NewMatrix([][]float64{
		{1, 2, 3, 4},
		{2, 4, 4, 2},
		{8, 6, 4, 1},
		{0, 0, 0, 1},
	})
	b := NewMatrix([][]float64{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	})

	tests := map[string]struct {
		m    *Matrix
		want *Matrix
	}{
		"simple": {m: a, want: NewMatrix([][]float64{
			{1, 2, 8, 0},
			{2, 4, 6, 0},
			{3, 4, 4, 0},
			{4, 2, 1, 1},
		})},
		"identity": {m: b, want: b},
	}

	for name, tc := range tests {
		got := Transpose(tc.m)

		if !IsEqual(got, tc.want) {
			t.Fatalf("%s: expected %v, got %v", name, tc.want, got)
		}
	}
}

func TestFindingDeterminant(t *testing.T) {
	a := NewMatrix([][]float64{{
		1, 5},
		{-3, 2},
	})
	b := NewMatrix([][]float64{
		{1, 2, 6},
		{-5, 8, -4},
		{2, 6, 4},
	})
	c := NewMatrix([][]float64{
		{-2, -8, 3, 5},
		{-3, 1, 7, 3},
		{1, 2, -9, 6},
		{-6, 7, 7, -9},
	})
	tests := map[string]struct {
		m    *Matrix
		want float64
	}{
		"2x2": {m: a, want: 17.0},
		"3x3": {m: b, want: -196.0},
		"4x4": {m: c, want: -4071},
	}

	for name, tc := range tests {
		got := GetDeterminant(tc.m)
		if got != tc.want {
			t.Fatalf("%s: expected %v, got %v", name, tc.want, got)
		}
	}
}

func TestSubmatrices(t *testing.T) {
	a := NewMatrix([][]float64{
		{1, 5, 0},
		{-3, 2, 7},
		{0, 6, -3},
	})
	b := NewMatrix([][]float64{
		{-6, 1, 1, 6},
		{-8, 5, 8, 6},
		{-1, 0, 8, 2},
		{-7, 1, -1, 1},
	})
	assert.True(t, IsEqual(GetSubmatrix(a, 0, 2), NewMatrix([][]float64{
		{-3, 2},
		{0, 6},
	})))
	assert.True(t, IsEqual(GetSubmatrix(b, 2, 1), NewMatrix([][]float64{
		{-6, 1, 6},
		{-8, 8, 6},
		{-7, -1, 1},
	})))
}

func TestMinors(t *testing.T) {
	a := NewMatrix([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	b := GetSubmatrix(a, 1, 0)
	assert.Equal(t, GetDeterminant(b), 25.0)
	assert.Equal(t, GetMinor(a, 1, 0), 25.0)
}

func TestCofactors(t *testing.T) {
	a := NewMatrix([][]float64{
		{3, 5, 0},
		{2, -1, -7},
		{6, -1, 5},
	})
	assert.Equal(t, GetMinor(a, 0, 0), -12.0)
	assert.Equal(t, GetCofactor(a, 0, 0), -12.0)
	assert.Equal(t, GetMinor(a, 1, 0), 25.0)
	assert.Equal(t, GetCofactor(a, 1, 0), -25.0)
}

func TestIsInvertible(t *testing.T) {
	a := NewMatrix([][]float64{
		{6, 4, 4, 4},
		{5, 5, 7, 6},
		{4, -9, 3, -7},
		{9, 1, 7, -6},
	})
	b := NewMatrix([][]float64{
		{-4, 2, -2, -3},
		{9, 6, 2, 6},
		{0, -5, 1, -5},
		{0, 0, 0, 0},
	})

	assert.True(t, IsInvertible(a))
	assert.False(t, IsInvertible(b))
}

func TestMatrixInversion(t *testing.T) {
	tests := []struct {
		input *Matrix
		want  *Matrix
	}{
		{
			input: NewMatrix([][]float64{
				{-5, 2, 6, -8},
				{1, -5, 1, 8},
				{7, 7, -6, -7},
				{1, -3, 7, 4},
			}),
			want: NewMatrix([][]float64{
				{0.21805, 0.45113, 0.24060, -0.04511},
				{-0.80827, -1.45677, -0.44361, 0.52068},
				{-0.07895, -0.22368, -0.05263, 0.19737},
				{-0.52256, -0.81391, -0.30075, 0.30639},
			}),
		},
		{
			input: NewMatrix([][]float64{
				{8, -5, 9, 2},
				{7, 5, 6, 1},
				{-6, 0, 9, 6},
				{-3, 0, -9, -4},
			}),
			want: NewMatrix([][]float64{
				{-0.15385, -0.15385, -0.28205, -0.53846},
				{-0.07692, 0.12308, 0.02564, 0.03077},
				{0.35897, 0.35897, 0.43590, 0.92308},
				{-0.69231, -0.69231, -0.76923, -1.92308},
			}),
		},
		{
			input: NewMatrix([][]float64{
				{9, 3, 0, 9},
				{-5, -2, -6, -3},
				{-4, 9, 6, 4},
				{-7, 6, 6, 2},
			}),
			want: NewMatrix([][]float64{
				{-0.04074, -0.07778, 0.14444, -0.22222},
				{-0.07778, 0.03333, 0.36667, -0.33333},
				{-0.02901, -0.14630, -0.10926, 0.12963},
				{0.17778, 0.06667, -0.26667, 0.33333},
			}),
		},
	}
	for _, tc := range tests {
		got, err := GetInverse(tc.input)
		if err != nil {
			t.Fatalf("expected error %v for input %v", err, tc.input)
		}
		if !IsEqual(got, tc.want) {
			t.Fatalf("expected %v, got %v", tc.want, got)
		}
	}
}

func TestMutliplyingByInverse(t *testing.T) {
	a := NewMatrix([][]float64{
		{3, -9, 7, 3},
		{3, -8, 2, -9},
		{-4, 4, 4, 1},
		{-6, 5, -1, 1},
	})
	b := NewMatrix([][]float64{
		{8, 2, 2, 2},
		{3, -1, 7, 0},
		{7, 0, 5, 4},
		{6, -2, 0, 5},
	})
	c, err := Multiply(a, b)
	assert.Nil(t, err)
	bInv, err := GetInverse(b)
	assert.Nil(t, err)
	d, err := Multiply(c, bInv)
	assert.Nil(t, err)
	assert.True(t, IsEqual(d, a))
}
