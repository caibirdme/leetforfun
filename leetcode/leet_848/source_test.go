package leet_848

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestShiftingLetters(t *testing.T) {
	var testCase = []struct {
		s      string
		shifts []int
		expect string
	}{
		{
			s:      "abc",
			shifts: []int{3, 5, 9},
			expect: "rpl",
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := shiftingLetters(tc.s, tc.shifts)
		ass.Equal(tc.expect, actual, "s: %s shifts: %v", tc.s, tc.shifts)
	}
}
