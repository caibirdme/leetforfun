package leet_4

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	if length%2 == 0 {
		t := length >> 1
		x1 := findKthElementInTwoSlice(nums1, nums2, t)
		x2 := findKthElementInTwoSlice(nums1, nums2, t+1)
		return float64(x1+x2) / 2
	} else {
		return float64(findKthElementInTwoSlice(nums1, nums2, (length+1)>>1))
	}
}

func findKthElementInTwoSlice(nums1, nums2 []int, k int) int {
	if nil == nums1 || 0 == len(nums1) {
		return findKthElementInOneSlice(nums2, k)
	} else if nil == nums2 || 0 == len(nums2) {
		return findKthElementInOneSlice(nums1, k)
	}
	leng1 := len(nums1)
	leng2 := len(nums2)
	mid1 := leng1 >> 1
	mid2 := leng2 >> 1
	switch {
	case nums1[mid1] <= nums2[mid2]:
		if k < mid1+mid2+2 {
			return findKthElementInTwoSlice(nums1, nums2[:mid2], k)
		} else {
			left := mid1 + 1
			if left >= leng1 {
				nums1 = nil
			} else {
				nums1 = nums1[left:]
			}
			return findKthElementInTwoSlice(nums1, nums2, k-mid1-1)
		}
	default:
		nums1, nums2 = nums2, nums1
		return findKthElementInTwoSlice(nums1, nums2, k)
	}
}

func findKthElementInOneSlice(nums []int, k int) int {
	return nums[k-1]
}
