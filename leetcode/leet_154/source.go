package leet_154

func findMin(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	if nums[0] == nums[length-1] {
		i := length - 1
		for ; i > 0 && nums[i] == nums[0]; i-- {
		}
		return solve(nums[:i+1])
	}
	return solve(nums)
}

func solve(nums []int) int {
	last := len(nums) - 1
	if last == 0 {
		return nums[0]
	}
	ans := nums[0]
	ll := nums[0]
	rr := nums[last]
	l, r := 0, last
	for l <= r {
		mid := (l + r) / 2
		t := nums[mid]
		if mid == r {
			return min(ans, t)
		}
		ans = min(ans, t)
		//right area
		if ll >= t && t <= rr {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
