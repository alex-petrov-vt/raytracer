package matrix

import (
	"errors"
	"testing"

	"github.com/alex-petrov-vt/raytracer/pkg/util"
)

func TestMatrices(t *testing.T) {
	m1 := NewMatrix([][]float64{{1, 2, 3, 4}, {5.5, 6.5, 7.5, 8.5}, {9, 10, 11, 12}, {13.5, 14.5, 15.5, 16.5}})
	m2 := NewMatrix([][]float64{{-3, 5}, {1, -2}})
	m3 := NewMatrix([][]float64{{-3, 5, 0}, {1, -2, -7}, {0, 1, 1}})
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
	a := NewMatrix([][]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}})
	b := NewMatrix([][]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}})
	c := NewMatrix([][]float64{{2, 3, 4, 5}, {6, 7, 8, 9}, {8, 7, 6, 5}, {4, 3, 2, 1}})
	d := NewMatrix([][]float64{{}})
	e := NewMatrix([][]float64{{}})
	f := NewMatrix([][]float64{{1, 2}, {3, 4}})
	g := NewMatrix([][]float64{{1, 2, 3, 4}, {5, 6, 7, 8}})
	h := NewMatrix([][]float64{{1}, {2}, {3}, {4}})

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
		got := Equals(tc.m1, tc.m2)
		if got != tc.want {
			t.Fatalf("%s: exptected %v, got %v", name, tc.want, got)
		}
	}
}

func TestMatrixMultiplication(t *testing.T) {
	a := NewMatrix([][]float64{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 8, 7, 6}, {5, 4, 3, 2}})
	b := NewMatrix([][]float64{{-2, 1, 2, 3}, {3, 2, 1, -1}, {4, 3, 6, 5}, {1, 2, 7, 8}})
	c := NewMatrix([][]float64{{20, 22, 50, 48}, {44, 54, 114, 108}, {40, 58, 110, 102}, {16, 26, 46, 42}})
	d := NewMatrix([][]float64{{}})
	e := NewMatrix([][]float64{{1, 2}, {3, 4}})

	tests := map[string]struct {
		m1   *Matrix
		m2   *Matrix
		want *Matrix
		err  error
	}{
		"simple":                  {m1: a, m2: b, want: c, err: nil},
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

		if !Equals(got, tc.want) {
			t.Fatalf("%s: expected %v, got %v", name, tc.want, got)
		}
	}

}
