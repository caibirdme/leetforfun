package leet_71

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSimplifyPath(t *testing.T) {
	var testCase = []struct {
		path   string
		expect string
	}{
		{
			path:   "/home/foo/..//bar//////////////",
			expect: "/home/bar",
		},
		{
			path:   "/home/foo/..",
			expect: "/home",
		},
		{
			path:   "/home/",
			expect: "/home",
		},
		{
			path:   "/a/./b/../../c/",
			expect: "/c",
		},
		{
			path:   "/../",
			expect: "/",
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		actual := simplifyPath(tc.path)
		should.Equal(tc.expect, actual, "path: %s", tc.path)
	}
}
