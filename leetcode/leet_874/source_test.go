package leet_874

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRobotSim(t *testing.T) {
	var testCase = []struct {
		commands  []int
		obstacles [][]int
		expect    int
	}{
		{
			commands: []int{1, 1, 3, 4, 3},
			obstacles: [][]int{
				[]int{-1, 5},
				[]int{-4, -4},
				[]int{-3, 3},
				[]int{3, 0},
				[]int{2, 5},
				[]int{-4, 4},
				[]int{-3, 1},
				[]int{-2, -4},
				[]int{-1, -4},
				[]int{0, -3},
			},
			expect: 144,
		},
		{
			commands:  []int{4, -1, 3},
			obstacles: nil,
			expect:    25,
		},
		{
			commands: []int{4, -1, 4, -2, 4},
			obstacles: [][]int{
				[]int{2, 4},
			},
			expect: 65,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		should.Equal(tc.expect, robotSim(tc.commands, tc.obstacles))
	}
}
