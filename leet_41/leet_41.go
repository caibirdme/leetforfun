package leet_41

func firstMissingPositive(nums []int) int {
	nums = append([]int{0}, nums...)
	length := len(nums)
	var t int
	for i := 0; i < length; i++ {
		t = nums[i]
		if t >= 0 && t < length {
			swap(nums, t)
		}
	}
	for i := 1; i < length; i++ {
		if nums[i] != i {
			return i
		}
	}
	return length
}

func swap(nums []int, i int) {
	length := len(nums)
	if i >= length {
		return
	}
	t := nums[i]
	if t == i {
		return
	}
	if t >= length || t <= 0 {
		nums[i] = i
		return
	}
	nums[i] = i
	swap(nums, t)
}
