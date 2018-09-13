package leet_162

func findPeakElement(nums []int) int {
	length := len(nums)
	if length == 1 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	last := length - 1
	if nums[last] > nums[last-1] {
		return last
	}
	for i := 1; i < last; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return last
}
