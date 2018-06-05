package leet_523

func checkSubarraySum(nums []int, k int) bool {
	length := len(nums)
	if length < 2 {
		return false
	}
	if k == 1 {
		return true
	}
	sum := make([]int, length+1)
	for i := 1; i <= length; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}
	if k == 0 {
		return checkZero(sum)
	}
	for i := 2; i <= length; i++ {
		for j := 0; j < i-1; j++ {
			if (sum[i]-sum[j])%k == 0 {
				return true
			}
		}
	}
	return false
}

func checkZero(sum []int) bool {
	length := len(sum)
	for i := 2; i < length; i++ {
		for j := 0; j < i-1; j++ {
			if (sum[i] - sum[j]) == 0 {
				return true
			}
		}
	}
	return false
}
