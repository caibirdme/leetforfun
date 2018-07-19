package leet_868

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinaryGap(t *testing.T) {
	var testCase = []struct {
		N      int
		expect int
	}{
		{
			N:      22,
			expect: 2,
		},
		{
			N:      5,
			expect: 2,
		},
		{
			N:      6,
			expect: 1,
		},
		{
			N:      8,
			expect: 0,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := binaryGap(tc.N)
		should.Equal(tc.expect, actual, "N:%d", tc.N)
	}
}
