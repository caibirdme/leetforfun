package leet_860

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLemonadeChange(t *testing.T) {
	var testCase = []struct {
		nums   []int
		expect bool
	}{
		{
			nums:   []int{5, 5, 5, 10, 5, 5, 10, 20, 20, 20},
			expect: false,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		should.Equal(tc.expect, lemonadeChange(tc.nums), "%v", tc.nums)
	}
}
