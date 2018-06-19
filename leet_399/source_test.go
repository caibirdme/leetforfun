package leet_399

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalcEquation(t *testing.T) {
	var testCase = []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		expect    []float64
	}{
		{
			equations: [][]string{
				[]string{"a", "b"},
				[]string{"a", "c"},
				[]string{"b", "e"},
				[]string{"c", "d"},
				[]string{"f", "e"},
			},
			values: []float64{2.0, 4.0, 3.0, 3.0, 0.5},
			queries: [][]string{
				[]string{"a", "b"},
				[]string{"b", "a"},
				[]string{"a", "e"},
				[]string{"a", "a"},
				[]string{"x", "x"},
				[]string{"a", "f"},
				[]string{"f", "d"},
				[]string{"c", "e"},
			},
			expect: []float64{2.0, 0.5, 6.0, 1.0, -1.0, 12.0, 1.0, 1.5},
		},
		{
			equations: [][]string{
				[]string{"a", "b"},
				[]string{"b", "c"},
				[]string{"bc", "cd"},
			},
			values: []float64{1.5, 2.5, 5.0},
			queries: [][]string{
				[]string{"a", "c"},
				[]string{"c", "b"},
				[]string{"bc", "cd"},
				[]string{"cd", "bc"},
			},
			expect: []float64{3.75, 0.4, 5.0, 0.2},
		},
		{
			equations: [][]string{
				[]string{"a", "b"},
				[]string{"b", "c"},
			},
			values: []float64{2.0, 3.0},
			queries: [][]string{
				[]string{"a", "c"},
				[]string{"b", "a"},
				[]string{"a", "e"},
				[]string{"a", "a"},
				[]string{"x", "x"},
			},
			expect: []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := calcEquation(tc.equations, tc.values, tc.queries)
		should.Equal(tc.expect, actual, "case#%d failed", idx)
	}
}
