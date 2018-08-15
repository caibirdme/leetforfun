package leet_130

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolve(t *testing.T) {
	var testCase = []struct {
		board  [][]byte
		expect [][]byte
	}{
		{
			board: [][]byte{
				[]byte("XXXX"),
				[]byte("XOOX"),
				[]byte("XXOX"),
				[]byte("XOXX"),
			},
			expect: [][]byte{
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XXXX"),
				[]byte("XOXX"),
			},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		solve(tc.board)
		should.Equal(tc.expect, tc.board, "case #%d fail", idx)
	}
}
