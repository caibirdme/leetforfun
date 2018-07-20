package leet_870

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAdvantageCount(t *testing.T) {
	var testCase = []struct {
		A, B   []int
		expect []int
	}{
		// {
		// 	A:      []int{2, 0, 4, 1, 2},
		// 	B:      []int{1, 3, 0, 0, 2},
		// 	expect: []int{2, 0, 2, 1, 4},
		// },
		{
			A:      []int{1, 3},
			B:      []int{2, 4},
			expect: []int{3, 1},
		},
		{
			A:      []int{2, 7, 11, 15},
			B:      []int{1, 10, 4, 11},
			expect: []int{2, 11, 7, 15},
		},
		{
			A:      []int{12, 24, 8, 32},
			B:      []int{13, 25, 32, 11},
			expect: []int{24, 32, 8, 12},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := advantageCount(tc.A, tc.B)
		should.Equal(tc.expect, actual, "case#%d fail", idx)
	}
}
