package leet_895

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFreqStack(t *testing.T) {
	type instruct struct {
		_type byte
		val   int
	}
	var testCase = []struct {
		instructions []instruct
		expect       []int
	}{
		{
			instructions: []instruct{
				{0, 5}, {0, 7}, {0, 5}, {0, 7}, {0, 4}, {0, 5}, {1, 0}, {1, 0}, {1, 0}, {1, 0},
			},
			expect: []int{5, 7, 5, 4},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		stack := Constructor()
		var actual []int
		for _, ins := range tc.instructions {
			_type, val := ins._type, ins.val
			if _type == 0 {
				stack.Push(val)
			} else {
				actual = append(actual, stack.Pop())
			}
		}
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}
