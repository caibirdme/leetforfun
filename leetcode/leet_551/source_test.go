package leet_551

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckRecord(t *testing.T) {
	var testCase = []struct {
		s      string
		expect bool
	}{
		{
			s:      "LLALL",
			expect: true,
		},
		{
			s:      "LLALLPLLPPPLPPPPPLLPLALPLL",
			expect: false,
		},
		{
			s:      "LLALLPLLPPPLPPPPPLLPLLPLL",
			expect: true,
		},
		{
			s:      "PPALLP",
			expect: true,
		},
		{
			s:      "PPALLL",
			expect: false,
		},
		{
			s:      "PPPPPPPPAA",
			expect: false,
		},
		{
			s:      "PPPPPPALLPPPPPPP",
			expect: true,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		should.Equal(tc.expect, checkRecord(tc.s), "s: %s", tc.s)
	}
}
