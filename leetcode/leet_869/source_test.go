package leet_869

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReorderedPowerOf2(t *testing.T) {
	var testCase = []struct {
		N      int
		expect bool
	}{
		{
			N:      1,
			expect: true,
		},
		{
			N:      10,
			expect: false,
		},
		{
			N:      16,
			expect: true,
		},
		{
			N:      24,
			expect: false,
		},
		{
			N:      46,
			expect: true,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := reorderedPowerOf2(tc.N)
		should.Equal(tc.expect, actual, "N:%d", tc.N)
	}
}
