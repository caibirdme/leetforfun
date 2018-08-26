package leet_892

func surfaceArea(grid [][]int) int {
	ans := 0
	n := len(grid)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			h := grid[i][j]
			// top
			if h == 0 {
				continue
			}
			ans += 2
			if i == 0 {
				ans += h
			} else {
				v := grid[i-1][j]
				if h > v {
					ans += h - v
				}
			}
			if j == 0 {
				ans += h
			} else {
				v := grid[i][j-1]
				if h > v {
					ans += h - v
				}
			}
			if i+1 == n {
				ans += h
			} else {
				v := grid[i+1][j]
				if h > v {
					ans += h - v
				}
			}
			if j+1 == n {
				ans += h
			} else {
				v := grid[i][j+1]
				if h > v {
					ans += h - v
				}
			}
		}
	}
	return ans
}
