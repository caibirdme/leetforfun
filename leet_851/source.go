package leet_851

func loudAndRich(richer [][]int, quiet []int) []int {
	lenq := len(quiet)
	connect := make([][]bool, lenq)
	for i := 0; i < lenq; i++ {
		connect[i] = make([]bool, lenq)
	}
	lenr := len(richer)
	for i := 0; i < lenr; i++ {
		rich, poor := richer[i][0], richer[i][1]
		connect[poor][rich] = true
	}
	dist := make([]int, lenq)
	for i := 0; i < lenq; i++ {
		dist[i] = -1
	}
	ans := make([]int, lenq)
	for i := 0; i < lenq; i++ {
		ans[i] = search(i, connect, quiet, dist)
	}
	return ans
}

func search(start int, connect [][]bool, quiet []int, dist []int) int {
	if dist[start] != -1 {
		return dist[start]
	}
	dist[start] = start
	length := len(quiet)
	for i := 0; i < length; i++ {
		if !connect[start][i] {
			continue
		}
		if quiet[dist[start]] > quiet[i] {
			dist[start] = i
		}
		t := search(i, connect, quiet, dist)
		if quiet[dist[start]] > quiet[t] {
			dist[start] = t
		}
	}
	if dist[start] == -1 {
		dist[start] = start
	}
	return dist[start]
}
