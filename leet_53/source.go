package leet_53

func maxSubArray(nums []int) int {
	currentMax := nums[0]
	ans := currentMax
	length := len(nums)
	for i := 1; i < length; i++ {
		t := currentMax + nums[i]
		if t > nums[i] {
			currentMax = t
		} else {
			currentMax = nums[i]
		}
		if currentMax > ans {
			ans = currentMax
		}
	}
	return ans
}
