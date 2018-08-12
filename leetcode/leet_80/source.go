package leet_80

func removeDuplicates(nums []int) int {
	length := len(nums)
	if length <= 2 {
		return length
	}
	count := 0
	i := 0
	for i < length {
		j := i + 1
		if j >= length {
			nums[count] = nums[i]
			count++
			break
		}
		if nums[i] != nums[j] {
			nums[count] = nums[i]
			count++
			i = j
			continue
		}
		k := j + 1
		for ; k < length && nums[k] == nums[j]; k++ {
		}
		nums[count] = nums[i]
		count++
		nums[count] = nums[j]
		count++
		i = k
	}
	return count
}
