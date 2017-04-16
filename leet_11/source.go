package leet_11

func maxArea(height []int) int {
	length := len(height)
	ans := 0
	for j := length - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if height[i] >= height[j] {
				area := height[j] * (j - i)
				if area > ans {
					ans = area
				}
				break
			} else {
				area := height[i] * (j - i)
				if area > ans {
					ans = area
				}
			}
		}
	}
	return ans
}
