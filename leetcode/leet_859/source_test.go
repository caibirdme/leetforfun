package leet_859

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBuddyStrings(t *testing.T) {
	var testCase = []struct {
		a, b   string
		expect bool
	}{
		{
			a:      "ab",
			b:      "ca",
			expect: false,
		},
		{
			a:      "",
			b:      "",
			expect: false,
		},
		{
			a:      "ab",
			b:      "ba",
			expect: true,
		},
		{
			a:      "ab",
			b:      "ab",
			expect: false,
		},
		{
			a:      "aa",
			b:      "aa",
			expect: true,
		},
		{
			a:      "aaaaaaabc",
			b:      "aaaaaaacb",
			expect: true,
		}, {
			a:      "",
			b:      "aa",
			expect: false,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := buddyStrings(tc.a, tc.b)
		should.Equal(tc.expect, actual, "A: %s | B: %s", tc.a, tc.b)
	}
}
