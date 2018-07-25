package leet_76

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinWindow(t *testing.T) {
	var testCase = []struct {
		s, t   string
		expect string
	}{
		{
			s:      "bba",
			t:      "a",
			expect: "a",
		},
		{
			s:      "ADOBECODEBANC",
			t:      "ABC",
			expect: "BANC",
		},
		{
			s:      "ADOBECODEBANCTTTTCAB",
			t:      "ABC",
			expect: "CAB",
		},
		{
			s:      "BAAAAAAAAAAAAAAAAATTTTTTTTTTTTTTTTTTTTTTTTC",
			t:      "ABC",
			expect: "BAAAAAAAAAAAAAAAAATTTTTTTTTTTTTTTTTTTTTTTTC",
		},
		{
			s:      "BAAAAAAAAAAAAAAAAATTTTTTTTTTTTTTTTTTTTTTTTC",
			t:      "ABCd",
			expect: "",
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := minWindow(tc.s, tc.t)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}
