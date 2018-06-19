package leet_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	data    []int
	target  int
	answers [2]int
}

func TestTwoSum(t *testing.T) {
	var data = []testCase{
		testCase{[]int{3, 2, 4}, 6, [2]int{1, 2}},
		testCase{[]int{2, 7, 11, 15}, 9, [2]int{0, 1}},
		testCase{[]int{7, 6, 5, 3, 2, 1, 4, 9, 10}, 17, [2]int{0, 8}},
	}
	ass := assert.New(t)
	for _, item := range data {
		a, b := twoSum(item.data, item.target)
		ass.Equal(item.answers[0], a)
		ass.Equal(item.answers[1], b)
	}
}
