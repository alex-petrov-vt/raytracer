package util

import (
	"testing"
)

func TestFloatEquals(t *testing.T) {
	tests := map[string]struct {
		input1 float64
		input2 float64
		want   bool
	}{
		"simple":                  {input1: (1.0 + 2.0), input2: 3.0, want: true},
		"zero":                    {input1: (3.153 - 3.153), input2: 0, want: true},
		"equal small numbers":     {input1: (0.005 * 0.005), input2: 0.000025, want: true},
		"not equal small numbers": {input1: (0.005 * 0.005), input2: 0.000026, want: false},
		"equal big numbers":       {input1: (1234.5 * 6789.0), input2: 8381020.5, want: true},
		"not equal big numbers":   {input1: (1234.5 * 6789.0), input2: 8381020.6, want: false},
	}

	for name, tc := range tests {
		got := FloatEquals(tc.input1, tc.input2)
		if got != tc.want {
			t.Fatalf("%s: expected: %v, got %v", name, tc.want, got)
		}
	}
}
