package leet_93

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRestoreIpAddresses(t *testing.T) {
	var testCase = []struct {
		s      string
		expect []string
	}{
		{
			s:      "255255255256",
			expect: nil,
		},
		{
			s:      "255255255255",
			expect: []string{"255.255.255.255"},
		},
		{
			s:      "8888",
			expect: []string{"8.8.8.8"},
		},
		{
			s:      "0000",
			expect: []string{"0.0.0.0"},
		},
		{
			s:      "19216811",
			expect: []string{"1.92.168.11", "19.2.168.11", "19.21.68.11", "19.216.8.11", "19.216.81.1", "192.1.68.11", "192.16.8.11", "192.16.81.1", "192.168.1.1"},
		},
		{
			s:      "25525511135",
			expect: []string{"255.255.11.135", "255.255.111.35"},
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := restoreIpAddresses(tc.s)
		should.Equal(len(tc.expect), len(actual), "length s: %s actual: %v", tc.s, actual)
		sort.Strings(tc.expect)
		sort.Strings(actual)
		should.Equal(tc.expect, actual, "s: %s", tc.s)
	}
}
