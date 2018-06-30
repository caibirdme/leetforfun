package leet_719

import (
	"sort"
)

func smallestDistancePair(nums []int, k int) int {
	length := len(nums)
	sort.Ints(nums)
	min := 0
	max := nums[length-1] - nums[0]
	ans := 0
	for min <= max {
		mid := (min + max) >> 1
		count := 0
		var left, right int
		for right = 1; right < length; right++ {
			for left = 0; left < right; left++ {
				if nums[right]-nums[left] <= mid {
					break
				}
			}
			count += right - left
			if count >= k {
				break
			}
		}
		if count >= k {
			ans = mid
			max = mid - 1
		} else {
			min = mid + 1
		}
	}
	return ans
}
