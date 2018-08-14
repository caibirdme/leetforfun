package kmp

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetNext(t *testing.T) {
	var testCase = []struct {
		sub    string
		expect []int
	}{
		{
			sub:    "abab",
			expect: []int{-1, -1, 0, 1},
		},
		{
			sub:    "abcdabcdeabcdabcdk",
			expect: []int{-1, -1, -1, -1, 0, 1, 2, 3, -1, 0, 1, 2, 3, 4, 5, 6, 7, -1},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := getNext([]byte(tc.sub))
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}

func TestFindSubStr(t *testing.T) {
	var testCase1 = []struct {
		s, sub string
		expect int
	}{
		{
			s:      "abcabsababff",
			sub:    "abcd",
			expect: -1,
		},
		{
			s:      "abcabsababff",
			sub:    "abab",
			expect: 6,
		},
	}
	should := require.New(t)
	for _, tc := range testCase1 {
		actual := FindSubStr(tc.s, tc.sub)
		should.Equal(tc.expect, actual, "s: %s sub: %s", tc.s, tc.sub)
	}
}

func TestFindSubStrWithStdLib(t *testing.T) {
	var testCase = []struct {
		s1, s2 string
	}{
		{
			s1: "abcabcabcasdaabcabcabcdqdashdjahs",
			s2: "abcabcabcd",
		},
		{
			s1: "abcabsababff",
			s2: "abcd",
		},
		{
			s1: "abcabsababff",
			s2: "abab",
		},
		{
			s1: "abcbdefghijklmn",
			s2: "abcbdefghijklmn",
		},
		{
			s1: "abcbdefghijklmnt",
			s2: "abcbdefghijklmnq",
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := FindSubStr(tc.s1, tc.s2)
		expect := strings.Index(tc.s1, tc.s2)
		should.Equal(expect, actual, "str: %s substr: %s", tc.s1, tc.s2)
	}
}
