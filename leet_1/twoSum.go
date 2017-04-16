package leet_1

import (
	"sort"
)

func twoSum(nums []int, target int) (int, int) {
	back := make([]int, len(nums))
	copy(back, nums)
	sort.Ints(nums)
	length := len(nums)
	var i, idx int
	for i = 0; i < length-1; i++ {
		t := target - nums[i]
		idx = sort.SearchInts(nums[i+1:], t)
		if i+idx+1 < length && nums[i+idx+1] == t {
			break
		}
	}
	a1, a2 := nums[i], nums[i+idx+1]
	var ansa, ansb int
	for j := 0; j < length; j++ {
		if back[j] == a1 || back[j] == a2 {
			if ansa == 0 {
				ansa = j
			} else {
				ansb = j
				break
			}
		}
	}
	if ansa > ansb {
		ansa, ansb = ansb, ansa
	}
	return ansa, ansb
}
