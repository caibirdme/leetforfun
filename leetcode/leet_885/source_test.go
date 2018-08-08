package leet_885

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumRescueBoats(t *testing.T) {
	var testCase = []struct {
		people []int
		limit  int
		expect int
	}{
		{
			people: []int{1, 3},
			limit:  3,
			expect: 2,
		},
		{
			people: []int{1, 2},
			limit:  3,
			expect: 1,
		},
		{
			people: []int{3, 2, 2, 1},
			limit:  3,
			expect: 3,
		},
		{
			people: []int{3, 5, 3, 4},
			limit:  5,
			expect: 4,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := numRescueBoats(tc.people, tc.limit)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}
