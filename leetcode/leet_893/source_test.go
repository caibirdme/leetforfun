package leet_893

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumSpecialEquivGroups(t *testing.T) {
	var testCase = []struct {
		A      []string
		expect int
	}{
		{
			A:      []string{"aa", "bb", "ab", "ba"},
			expect: 4,
		},
		{
			A:      []string{"a", "b", "c", "a", "c", "c"},
			expect: 3,
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := numSpecialEquivGroups(tc.A)
		should.Equal(tc.expect, actual, "case #%d fail", idx)
	}
}
