package leet_818

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRacecar(t *testing.T) {
	ass := require.New(t)
	var testCase = []int{4, 11, 0, 1, 2, 5, 10, 50, 100, 200, 500, 1000, 1024, 1023, 5000, 10000}
	for _, target := range testCase {
		expect := racecar_bfs(target)
		actual := racecar(target)
		ass.Equal(expect, actual, "target: %d", target)
	}
}
