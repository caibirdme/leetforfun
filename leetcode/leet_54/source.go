package leet_54

func spiralOrder(matrix [][]int) []int {
	xstep := len(matrix)
	if 0 == xstep {
		return nil
	}
	ystep := len(matrix[0])
	if 0 == ystep {
		return nil
	}
	ans := make([]int, 0, xstep*ystep)
	right(matrix, point{0, 0}, xstep, ystep, &ans)
	return ans
}

type point struct {
	x, y int
}

func up(matrix [][]int, now point, xstep, ystep int, ans *[]int) {
	for i := 0; i < xstep; i++ {
		*ans = append(*ans, matrix[now.x+i][now.y])
	}
	if ystep-1 == 0 {
		return
	}
	left(matrix, point{x: now.x + xstep - 1, y: now.y - 1}, xstep, ystep-1, ans)
}

func down(matrix [][]int, now point, xstep, ystep int, ans *[]int) {
	for i := 0; i < xstep; i++ {
		*ans = append(*ans, matrix[now.x-i][now.y])
	}
	if ystep-1 == 0 {
		return
	}
	right(matrix, point{x: now.x - xstep + 1, y: now.y + 1}, xstep, ystep-1, ans)
}

func left(matrix [][]int, now point, xstep, ystep int, ans *[]int) {
	for i := 0; i < ystep; i++ {
		*ans = append(*ans, matrix[now.x][now.y-i])
	}
	if xstep-1 == 0 {
		return
	}
	down(matrix, point{x: now.x - 1, y: now.y - ystep + 1}, xstep-1, ystep, ans)
}

func right(matrix [][]int, now point, xstep, ystep int, ans *[]int) {
	for i := 0; i < ystep; i++ {
		*ans = append(*ans, matrix[now.x][now.y+i])
	}
	if xstep-1 == 0 {
		return
	}
	up(matrix, point{x: now.x + 1, y: now.y + ystep - 1}, xstep-1, ystep, ans)
}
