package leet_849

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMaxDistToClosest(t *testing.T) {
	var testCase = []struct {
		seats  []int
		expect int
	}{
		{
			seats:  []int{1, 0, 0, 0, 1, 0, 1},
			expect: 2,
		},
		{
			seats:  []int{1, 0, 0, 0},
			expect: 3,
		},
		{
			seats:  []int{0, 0, 0, 0, 1},
			expect: 4,
		},
		{
			seats:  []int{0, 0, 0, 0, 1, 0, 0, 0},
			expect: 4,
		},
		{
			seats:  []int{1, 0, 0, 0, 0, 1, 0, 0, 0},
			expect: 3,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := maxDistToClosest(tc.seats)
		ass.Equal(tc.expect, actual, "seats: %v", tc.seats)
	}
}
