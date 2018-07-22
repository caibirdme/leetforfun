package leet_643

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	length := len(nums)
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	ans := float64(sum) / float64(k)
	left := 0
	for i := k; i < length; i++ {
		sum = sum - nums[left] + nums[i]
		ans = max(ans, float64(sum)/float64(k))
		left++
	}
	return ans
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
