package qsort

func Sort(nums []int) {
	length := len(nums) - 1
	quickSort(nums, 0, length)
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	p := partition(nums, left, right)
	quickSort(nums, left, p-1)
	quickSort(nums, p+1, right)
}

func partition(nums []int, left, right int) int {
	mid := (left + right) >> 1
	pivot := nums[mid]
	nums[left], nums[mid] = pivot, nums[left]
	i := left
	j := right
	for i < j {
		for nums[j] >= pivot && i < j {
			j--
		}
		for nums[i] <= pivot && i < j {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[left], nums[j] = nums[j], pivot
	return j
}
