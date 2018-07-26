package leet_115

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumDistinct(t *testing.T) {
	var testCase = []struct {
		s, t   string
		expect int
	}{
		{
			s:      "",
			t:      "a",
			expect: 0,
		},
		{
			s:      "aaaaaabb",
			t:      "aaab",
			expect: 40, // C63*C21
		},
		{
			s:      "aaaaaa",
			t:      "aaa",
			expect: 20, // C63
		},
		{
			s:      "iiiii",
			t:      "ii",
			expect: 10, // C52
		},
		{
			s:      "iiiii",
			t:      "i",
			expect: 5,
		},
		{
			s:      "rabbbit",
			t:      "rabbit",
			expect: 3,
		},
		{
			s:      "babgbag",
			t:      "bag",
			expect: 5,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := numDistinct(tc.s, tc.t)
		actual_2 := numDistinct2(tc.s, tc.t)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
		should.Equal(tc.expect, actual_2, "case #%d actual_2 fail", idx)

	}
}
