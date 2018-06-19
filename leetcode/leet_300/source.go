package leet_300

func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	f := make([]int, length)
	ans := 0
	for i := 1; i < length; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				if t := f[j] + 1; t > f[i] {
					f[i] = t
					ans = max(ans, t)
				}
			}
		}
	}
	return ans + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
