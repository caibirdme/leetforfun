package leet_69

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMyqsrt(t *testing.T) {
	var testCase = []struct {
		x      int
		expect int
	}{
		{
			x:      2147395600,
			expect: 10,
		},
		{
			x:      8,
			expect: 2,
		},
		{
			x:      16,
			expect: 4,
		},
		{
			x:      25,
			expect: 5,
		},
		{
			x:      26,
			expect: 5,
		},
		{
			x:      90,
			expect: 9,
		},
		{
			x:      0,
			expect: 0,
		},
		{
			x:      1,
			expect: 1,
		},
		{
			x:      2,
			expect: 1,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := mySqrt(tc.x)
		should.Equal(tc.expect, actual, "x: %d", tc.x)
	}
}
