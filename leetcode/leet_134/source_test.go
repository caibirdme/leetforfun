package leet_134

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCompleteCircuit(t *testing.T) {
	var testCase = []struct {
		gas    []int
		cost   []int
		expect int
	}{
		{
			gas:    []int{1, 2, 3, 4, 5},
			cost:   []int{3, 4, 5, 1, 2},
			expect: 3,
		},
		{
			gas:    []int{2, 3, 4},
			cost:   []int{3, 4, 3},
			expect: -1,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := canCompleteCircuit(tc.gas, tc.cost)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}
