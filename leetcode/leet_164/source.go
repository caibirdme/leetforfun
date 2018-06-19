package leet_164

func maximumGap(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	minNumber, maxNumber := findMin(nums), findMax(nums)
	if minNumber == maxNumber {
		return 0
	}
	dist := (maxNumber - minNumber) / (len(nums) - 1)
	if (maxNumber-minNumber)%(len(nums)-1) != 0 {
		dist++
	}
	bucketLen := (maxNumber-minNumber)/dist + 1
	arr := make([]*bucket, bucketLen)
	for idx := range arr {
		arr[idx] = new(bucket)
	}
	for _, n := range nums {
		idx := (n - minNumber) / dist
		arr[idx].Insert(n)
	}
	var diff int
	for i := 1; i < bucketLen; i++ {
		if !arr[i].setted {
			arr[i].Insert(arr[i-1].GetMax())
		} else {
			t := arr[i].GetMin() - arr[i-1].GetMax()
			diff = max(diff, t)
		}
	}
	return diff
}

func findMin(nums []int) int {
	minNumber := nums[0]
	for i := 1; i < len(nums); i++ {
		if minNumber > nums[i] {
			minNumber = nums[i]
		}
	}
	return minNumber
}

func findMax(nums []int) int {
	maxNumber := nums[0]
	for i := 1; i < len(nums); i++ {
		if maxNumber < nums[i] {
			maxNumber = nums[i]
		}
	}
	return maxNumber
}

type bucket struct {
	setted   bool
	max, min int
}

func (b *bucket) Insert(x int) {
	if !b.setted {
		b.setted = true
		b.max, b.min = x, x
		return
	}
	if b.max < x {
		b.max = x
	}
	if b.min > x {
		b.min = x
	}
}

func (b *bucket) GetMax() int {
	return b.max
}

func (b *bucket) GetMin() int {
	return b.min
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
