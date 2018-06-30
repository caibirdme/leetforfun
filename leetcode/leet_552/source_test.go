package leet_552

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckRecord(t *testing.T) {
	var testCase = []struct {
		n int
		m int
	}{
		{
			n: 1,
			m: 3,
		},
		{
			n: 3,
			m: 19,
		},
		{
			n: 2,
			m: 8,
		},
	}
	should := require.New(t)
	for _, tc := range testCase {
		should.Equal(tc.m, checkRecord(tc.n), "n: %d", tc.n)
	}
}
