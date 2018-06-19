package leet_560

func subarraySum(nums []int, k int) int {
	length := len(nums)
	sum := make([]int, length+1)
	count := 0
	for i := 0; i < length; i++ {
		sum[i+1] = sum[i] + nums[i]
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j <= length; j++ {
			if sum[j]-sum[i] == k {
				count++
			}
		}
	}
	return count
}
