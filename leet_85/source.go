package leet_85

func maximalRectangle(matrix [][]byte) int {
	n := len(matrix)
	if n == 0 {
		return 0
	}
	m := len(matrix[0])
	if m == 0 {
		return 0
	}
	successive := make([][]int, n)
	for i := 0; i < n; i++ {
		successive[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		successive[i][m-1] = int(matrix[i][m-1] - '0')
		for j := m - 2; j >= 0; j-- {
			if matrix[i][j] == '1' {
				successive[i][j] = successive[i][j+1] + 1
			}
		}
	}
	area := 0
	for i := 0; i < n; i++ {
		maxHeight := n - i
		for j := 0; j < m; j++ {
			if successive[i][j] == 0 {
				continue
			}
			width := successive[i][j]
			area = max(width, area)
			for height := 2; height <= maxHeight; height++ {
				v := successive[i+height-1][j]
				if v == 0 {
					break
				}
				width = min(v, width)
				t := height * width
				area = max(area, t)
			}
		}
	}
	return area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
