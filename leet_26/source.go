package leet_26

func removeDuplicates(nums []int) int {
	if nil == nums || 0 == len(nums) {
		return 0
	}
	var ans int
	pre := nums[0]
	length := len(nums)
	for i := 1; i < length; i++ {
		if nums[i] == pre {
			continue
		}
		pre = nums[i]
		ans++
		nums[ans] = pre
	}
	return ans + 1
}
