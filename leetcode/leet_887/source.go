package leet_887

func projectionArea(grid [][]int) int {
	n := len(grid)
	ans := 0
	yz := make([]int, n)
	xz := make([]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			v := grid[i][j]
			if v == 0 {
				continue
			}
			ans++
			yz[j] = max(yz[j], v)
			xz[i] = max(xz[i], v)
		}
	}
	for i := 0; i < n; i++ {
		ans += yz[i] + xz[i]
	}
	return ans
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
