package leet_64

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	if n == 0 || m == 0 {
		return 0
	}
	f := make([][]int, n)
	for i := 0; i < n; i++ {
		f[i] = make([]int, m)
		for j := 0; j < m; j++ {
			f[i][j] = -1
		}
	}
	f[0][0] = grid[0][0]
	return search(len(grid)-1, len(grid[0])-1, grid, f)
}

func search(i, j int, grid, f [][]int) int {
	if f[i][j] > -1 {
		return f[i][j]
	}
	if i == 0 {
		f[0][j] = search(0, j-1, grid, f) + grid[0][j]
		return f[0][j]
	}
	if j == 0 {
		f[i][0] = search(i-1, 0, grid, f) + grid[i][0]
		return f[i][0]
	}
	f[i][j] = min(search(i-1, j, grid, f), search(i, j-1, grid, f)) + grid[i][j]
	return f[i][j]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
