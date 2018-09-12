package leet_148

import (
	"sort"
	"testing"

	"github.com/caibirdme/leetforfun/testsuit"
	"github.com/stretchr/testify/require"
)

func TestSortList(t *testing.T) {
	var testCase = []struct {
		nums []int
	}{
		{
			nums: []int{1, 2},
		},
		{
			nums: []int{3, 1, 2},
		},
		{
			nums: []int{3, 1, 4, 2, 2},
		},
		{
			nums: []int{},
		},
		{
			nums: []int{1},
		},
		{
			nums: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0, -1},
		},
		{
			nums: []int{2, 1, 4, 3, 6, 5, 8, 7, 9, 10, 12, 11},
		},
	}
	should := require.New(t)
	for idx, tc := range testCase {
		actual := sortList(testsuit.CreateList(tc.nums))
		sort.Ints(tc.nums)
		expect := testsuit.CreateList(tc.nums)
		should.True(testsuit.CmpList(expect, actual), "case #%d fail", idx)
	}
}
