package leet_97

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsInterleave(t *testing.T) {
	var testCase = []struct {
		s1, s2, s3 string
		expect     bool
	}{
		{
			s1:     "db",
			s2:     "b",
			s3:     "cbb",
			expect: false,
		},
		{
			s1:     "db",
			s2:     "c",
			s3:     "cbb",
			expect: false,
		},
		{
			s1:     "bb",
			s2:     "c",
			s3:     "cbb",
			expect: true,
		},
		{
			s1:     "",
			s2:     "",
			s3:     "",
			expect: true,
		},
		{
			s1:     "ab",
			s2:     "cd",
			s3:     "acbd",
			expect: true,
		},
		{
			s1:     "ab",
			s2:     "cd",
			s3:     "abcde",
			expect: false,
		},
		{
			s1:     "aabcc",
			s2:     "dbbca",
			s3:     "aadbbcbcac",
			expect: true,
		},
		{
			s1:     "aabcc",
			s2:     "dbbca",
			s3:     "aadbbbaccc",
			expect: false,
		},
		{
			s1:     "aabcc",
			s2:     "",
			s3:     "aabcc",
			expect: true,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := isInterleave(tc.s1, tc.s2, tc.s3)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}
