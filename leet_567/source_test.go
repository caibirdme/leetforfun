package leet_567

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckInclusion(t *testing.T) {
	var testCase = []struct {
		s1, s2 string
		expect bool
	}{
		{
			s1:     "ab",
			s2:     "bac",
			expect: true,
		},
		{
			s1:     "ab",
			s2:     "eidbaooo",
			expect: true,
		},
		{
			s1:     "ab",
			s2:     "eidboaoo",
			expect: false,
		},
	}
	ass := require.New(t)
	for _, tc := range testCase {
		actual := checkInclusion(tc.s1, tc.s2)
		ass.Equal(tc.expect, actual, "s1:%s s2:%s", tc.s1, tc.s2)
	}
}
