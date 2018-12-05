package leet_189

func rotate(nums []int, k int) {
	length := len(nums)
	k %= length
	if k == 0 {
		return
	}
	shift := make([]int, k)
	last := length - k
	copy(shift, nums[last:])
	for i := last - 1; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for i := 0; i < k; i++ {
		nums[i] = shift[i]
	}
}
