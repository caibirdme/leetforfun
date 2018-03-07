package leet_56

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQsort(t *testing.T) {
	var testData = []struct {
		in  []Interval
		out []Interval
	}{
		{
			in: []Interval{
				Interval{Start: 1, End: 4},
				Interval{Start: 1, End: 4},
			},
			out: []Interval{
				Interval{Start: 1, End: 4},
				Interval{Start: 1, End: 4},
			},
		},
		{
			in: []Interval{
				Interval{Start: 5, End: 8},
				Interval{Start: 4, End: 9},
				Interval{Start: 3, End: 4},
				Interval{Start: 2, End: 8},
				Interval{Start: 2, End: 7},
				Interval{Start: 1, End: 6},
			},
			out: []Interval{
				Interval{Start: 1, End: 6},
				Interval{Start: 2, End: 7},
				Interval{Start: 2, End: 8},
				Interval{Start: 3, End: 4},
				Interval{Start: 4, End: 9},
				Interval{Start: 5, End: 8},
			},
		},
		{
			in: []Interval{
				Interval{Start: 5, End: 8},
				Interval{Start: 6, End: 9},
				Interval{Start: 3, End: 4},
				Interval{Start: 0, End: 7},
				Interval{Start: 5, End: 6},
			},
			out: []Interval{
				Interval{Start: 0, End: 7},
				Interval{Start: 3, End: 4},
				Interval{Start: 5, End: 6},
				Interval{Start: 5, End: 8},
				Interval{Start: 6, End: 9},
			},
		},
		{
			in: []Interval{
				Interval{Start: 5, End: 8},
				Interval{Start: 4, End: 9},
				Interval{Start: 3, End: 4},
				Interval{Start: 2, End: 7},
				Interval{Start: 1, End: 6},
			},
			out: []Interval{
				Interval{Start: 1, End: 6},
				Interval{Start: 2, End: 7},
				Interval{Start: 3, End: 4},
				Interval{Start: 4, End: 9},
				Interval{Start: 5, End: 8},
			},
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		quickSort(tc.in, 0, len(tc.in)-1)
		ass.Equal(tc.out, tc.in)
	}
}

func TestMerge(t *testing.T) {
	var testData = []struct {
		in  []Interval
		out []Interval
	}{
		{
			in: []Interval{
				Interval{Start: 1, End: 3},
				Interval{Start: 2, End: 6},
				Interval{Start: 8, End: 10},
				Interval{Start: 15, End: 18},
			},
			out: []Interval{
				Interval{Start: 1, End: 6},
				Interval{Start: 8, End: 10},
				Interval{Start: 15, End: 18},
			},
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, merge(tc.in))
	}
}
