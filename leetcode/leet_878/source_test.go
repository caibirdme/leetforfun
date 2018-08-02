package leet_878

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNthMagicalNumber(t *testing.T) {
	var testCase = []struct {
		N, A, B int
		expect  int
	}{
		{
			N:      4,
			A:      2,
			B:      3,
			expect: 6,
		},
		{
			N:      1000000000,
			A:      39999,
			B:      40000,
			expect: 999860007,
		},
		{
			N:      10,
			A:      10,
			B:      8,
			expect: 50,
		},
		{
			N:      1,
			A:      2,
			B:      3,
			expect: 2,
		},
		{
			N:      5,
			A:      2,
			B:      4,
			expect: 10,
		},
		{
			N:      3,
			A:      6,
			B:      4,
			expect: 8,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := nthMagicalNumber(tc.N, tc.A, tc.B)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}
