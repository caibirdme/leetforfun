package leet_885

import (
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	length := len(people)
	if length <= 1 {
		return length
	}
	sort.Ints(people)
	ans := 0
	r := length - 1
	for r >= 0 && people[r] == limit {
		ans++
		r--
	}
	i := 0
	for i < r {
		if people[i]+people[r] <= limit {
			ans++
			i++
			r--
		} else {
			ans++
			r--
		}
	}
	if i == r {
		ans++
	}
	return ans
}
