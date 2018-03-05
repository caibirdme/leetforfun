package leet_55

func canJump(nums []int) bool {
	length := len(nums)
	if length == 1 {
		return true
	}
	if nums[0] == 0 {
		return false
	}
	farthest := 0
	l := length - 1
	for i := 0; i < l; i++ {
		if i > farthest {
			return false
		}
		t := i + nums[i]
		if t > farthest {
			farthest = t
		}
		if farthest >= l {
			return true
		}
	}
	return false
}
