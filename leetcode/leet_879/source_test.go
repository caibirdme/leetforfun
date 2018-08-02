package leet_879

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProfitableSchemes(t *testing.T) {
	var testCase = []struct {
		G, P          int
		group, profit []int
		expect        int
	}{
		{
			G:      3,
			P:      1,
			group:  []int{1, 1, 1},
			profit: []int{1, 1, 1},
			expect: 7,
		},
		{
			G:      3,
			P:      3,
			group:  []int{3},
			profit: []int{5},
			expect: 1,
		},
		{
			G:      5,
			P:      3,
			group:  []int{2, 2},
			profit: []int{2, 3},
			expect: 2,
		},
		{
			G:      10,
			P:      5,
			group:  []int{2, 3, 5},
			profit: []int{6, 7, 8},
			expect: 7,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := profitableSchemes(tc.G, tc.P, tc.group, tc.profit)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}
