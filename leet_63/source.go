package leet_63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	length := len(obstacleGrid)
	mem := make([][]int, length)
	for i := 0; i < length; i++ {
		mem[i] = make([]int, len(obstacleGrid[i]))
	}
	mem[0][0] = 1
	x := length - 1
	y := len(mem[0]) - 1
	return calc(x, y, mem, obstacleGrid)
}

func calc(x, y int, mem, obstacleGrid [][]int) int {
	if x < 0 || y < 0 {
		return 0
	}
	if obstacleGrid[x][y] == 1 {
		return 0
	}
	if mem[x][y] != 0 {
		return mem[x][y]
	}
	mem[x][y] = calc(x-1, y, mem, obstacleGrid) + calc(x, y-1, mem, obstacleGrid)
	return mem[x][y]
}
