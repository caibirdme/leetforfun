package leet_209

func minSubArrayLen(s int, nums []int) int {
	sum := 0
	count := 0
	length := len(nums)
	i := 0
	for ; i < length; i++ {
		sum += nums[i]
		count++
		if sum >= s {
			i++
			break
		}
	}
	if i == length && sum < s {
		return 0
	}
	ans := count
	head := 0
	for sum >= s && head < i {
		t := sum - nums[head]
		if t < s {
			break
		}
		sum = t
		head++
		count--
	}
	if ans > count {
		ans = count
	}
	for ; i < length; i++ {
		sum += nums[i] - nums[head]
		head++
		if head == i {
			continue
		}
		for sum >= s && head < i {
			t := sum - nums[head]
			if t < s {
				break
			}
			sum = t
			head++
			count--
		}
		if ans > count {
			ans = count
		}
	}
	return ans
}
