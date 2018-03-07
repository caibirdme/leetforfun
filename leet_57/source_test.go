package leet_57

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	var testData = []struct {
		in   []Interval
		newI Interval
		out  []Interval
	}{
		{
			in: []Interval{
				Interval{Start: 3, End: 5},
				Interval{Start: 12, End: 15},
			},
			newI: Interval{Start: 6, End: 6},
			out: []Interval{
				Interval{Start: 3, End: 5},
				Interval{Start: 6, End: 6},
				Interval{Start: 12, End: 15},
			},
		},
		{
			in: []Interval{
				Interval{Start: 1, End: 5},
			},
			newI: Interval{Start: 6, End: 8},
			out: []Interval{
				Interval{Start: 1, End: 5},
				Interval{Start: 6, End: 8},
			},
		},
		{
			in: []Interval{
				Interval{Start: 1, End: 3},
			},
			newI: Interval{Start: 2, End: 5},
			out: []Interval{
				Interval{Start: 1, End: 5},
			},
		},
		{
			in: []Interval{
				Interval{Start: 1, End: 3},
				Interval{Start: 6, End: 9},
			},
			newI: Interval{Start: 2, End: 5},
			out: []Interval{
				Interval{Start: 1, End: 5},
				Interval{Start: 6, End: 9},
			},
		},
		{
			in: []Interval{
				Interval{Start: 1, End: 2},
				Interval{Start: 3, End: 5},
				Interval{Start: 6, End: 7},
				Interval{Start: 8, End: 10},
				Interval{Start: 12, End: 16},
			},
			newI: Interval{Start: 4, End: 9},
			out: []Interval{
				Interval{Start: 1, End: 2},
				Interval{Start: 3, End: 10},
				Interval{Start: 12, End: 16},
			},
		},
	}
	ass := assert.New(t)
	for _, tc := range testData {
		ass.Equal(tc.out, insert(tc.in, tc.newI))
	}
}
