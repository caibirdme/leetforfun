package leet_884

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecodeAtIndex(t *testing.T) {
	var testCase = []struct {
		S      string
		K      int
		expect string
	}{
		{
			S:      "vk6u5xhq9v",
			K:      554,
			expect: "k",
		},
		{
			S:      "ab23c2def",
			K:      27,
			expect: "d",
		},
		{
			S:      "ab23c2def",
			K:      28,
			expect: "e",
		},
		{
			S:      "ab23c2def",
			K:      29,
			expect: "f",
		},
		{
			S:      "ab23c2",
			K:      7,
			expect: "a",
		},
		{
			S:      "ab23c2",
			K:      13,
			expect: "c",
		},
		{
			S:      "ab23c2",
			K:      14,
			expect: "a",
		},
		{
			S:      "ab23c2",
			K:      15,
			expect: "b",
		},
		{
			S:      "ab23c2",
			K:      26,
			expect: "c",
		},
		{
			S:      "ab23c2",
			K:      12,
			expect: "b",
		},
		{
			S:      "b2a2c2d2",
			K:      16,
			expect: "b",
		},
		{
			S:      "leet2code3",
			K:      25,
			expect: "l",
		},
		{
			S:      "leet2code3",
			K:      28,
			expect: "t",
		},
		{
			S:      "leet2code3",
			K:      10,
			expect: "o",
		},
		{
			S:      "leet2code3",
			K:      12,
			expect: "e",
		},
		{
			S:      "ha22",
			K:      5,
			expect: "h",
		},
		{
			S:      "a2345678999999999999999",
			K:      1,
			expect: "a",
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := decodeAtIndex(tc.S, tc.K)
		should.Equal(tc.expect, actual, "S: %s K: %d", tc.S, tc.K)
	}
}
