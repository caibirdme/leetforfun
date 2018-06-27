package leet_546

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRemoveBoxes(t *testing.T) {
	should := require.New(t)
	// expect := removeBoxes2([]int{5, 8, 2, 7, 5, 7, 4, 10, 5, 5, 8, 8, 7, 8})
	// actual := removeBoxes([]int{5, 8, 2, 7, 5, 7, 4, 10, 5, 5, 8, 8, 7, 8})
	// should.Equal(expect, actual)
	const n = 30
	boxes := make([]int, n)
	for i := 0; i < 1000; i++ {
		for j := 0; j < n; j++ {
			boxes[j] = rand.Intn(10) + 1
		}
		expect := removeBoxes2(boxes)
		actual := removeBoxes(boxes)
		should.Equal(expect, actual, "boxes: %v", boxes)
	}
}
