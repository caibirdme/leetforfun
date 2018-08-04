package leet_59

func generateMatrix(n int) [][]int {
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		ans[i] = make([]int, n)
	}
	var step = [4][2]int{
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{0, -1},
		[2]int{-1, 0},
	}
	x, y := 0, -1
	count := 1
	t := 1
	for n > 0 {
		for j := 0; j < 4; j++ {
			t++
			if t > 2 {
				t = 1
				n--
			}
			for i := 0; i < n; i++ {
				x += step[j][0]
				y += step[j][1]
				ans[x][y] = count
				count++
			}
		}
	}
	return ans
}
