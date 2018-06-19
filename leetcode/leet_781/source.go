package leet_781

import (
	"sort"
)

func numRabbits(answers []int) int {
	length := len(answers)
	if 0 == length {
		return 0
	}
	if 1 == length {
		return answers[0] + 1
	}
	sort.Ints(answers)
	colors := 0
	for i := 0; i < length; i++ {
		num := answers[i]
		colors += num + 1
		var j int
		for j = 1; j <= num; j++ {
			if t := i + j; t >= length || answers[t] != num {
				break
			}
		}
		i += j - 1
	}
	return colors
}
