package leet_902

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtMostNGivenDigitSet(t *testing.T) {
	var testCase = []struct {
		D      []string
		N      int
		expect int
	}{
		{
			D:      []string{"1", "4", "9"},
			N:      1000000000,
			expect: 29523,
		},
		{
			D:      []string{"1", "3", "5", "7"},
			N:      100,
			expect: 20,
		},
		{
			D:      []string{"3", "4", "5", "6"},
			N:      64,
			expect: 18,
		},

		{
			D:      []string{"2", "3", "5", "7"},
			N:      3,
			expect: 2,
		},
		{
			D:      []string{"2", "3", "5", "7"},
			N:      1,
			expect: 0,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := atMostNGivenDigitSet(tc.D, tc.N)
		should.Equal(tc.expect, actual, "D: %v N: %d", tc.D, tc.N)
	}
}
