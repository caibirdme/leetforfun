package leet_96

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumTrees(t *testing.T) {
	var testCase = []struct {
		n      int
		expect int
	}{
		{
			n:      1,
			expect: 1,
		},
		{
			n:      3,
			expect: 5,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := numTrees(tc.n)
		should.Equal(tc.expect, actual, "n: %d", tc.n)
	}
}
