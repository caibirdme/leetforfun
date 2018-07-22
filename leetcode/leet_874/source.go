package leet_874

import (
	"fmt"
)

func robotSim(commands []int, obstacles [][]int) int {
	var steps = [...][2]int{
		[2]int{0, 1},
		[2]int{1, 0},
		[2]int{0, -1},
		[2]int{-1, 0},
	}
	ans := 0
	obs := make(map[string]struct{})
	for _, obstacle := range obstacles {
		x, y := obstacle[0], obstacle[1]
		obs[fmt.Sprintf("%d,%d", x, y)] = struct{}{}
	}
	direction := 0
	x, y := 0, 0
	for _, command := range commands {
		if command == -1 {
			direction = (direction + 1) % 4
		} else if command == -2 {
			direction = (direction + 3) % 4
		} else {
			dx, dy := steps[direction][0], steps[direction][1]
			for i := 0; i < command; i++ {
				tx, ty := x+dx, y+dy
				if _, ok := obs[fmt.Sprintf("%d,%d", tx, ty)]; ok {
					break
				}
				x, y = tx, ty
			}
			ans = max(ans, x*x+y*y)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
