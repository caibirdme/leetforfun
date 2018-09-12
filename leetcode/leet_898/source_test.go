package leet_898

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSubarrayBitwiseORs(t *testing.T) {
	var testCase = []struct {
		A      []int
		expect int
	}{
		{
			A:      []int{1, 2, 4},
			expect: 6,
		},
		{
			A:      []int{1, 1, 2},
			expect: 3,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := subarrayBitwiseORs(tc.A)
		should.Equal(tc.expect, actual, "A: %v", tc.A)
	}
}
