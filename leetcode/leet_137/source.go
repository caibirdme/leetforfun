package leet_137

func singleNumber(nums []int) int {
	ans := uint32(0)
	for idx := uint32(0); idx < 32; idx++ {
		count := uint32(0)
		for _, n := range nums {
			if uint32(n)&(1<<idx) != 0 {
				count++
			}
		}
		if count%3 != 0 {
			ans |= 1 << idx
		}
	}
	return int(int32(ans))
}
