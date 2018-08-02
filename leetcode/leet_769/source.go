package leet_769

func maxChunksToSorted(arr []int) int {
	length := len(arr) - 1
	if length == 0 {
		return 1
	}
	var visited [10]bool
	count := 1
	i := 0
	for i <= length {
		idx := findIdx(arr, i, &visited)
		if idx >= length {
			return count
		}
		right := idx
		for j := i + 1; j <= right; j++ {
			if !visited[j] {
				right = findIdx(arr, j, &visited)
			}
		}
		if right < length {
			count++
		}
		i = right + 1
	}
	return count
}

func findIdx(arr []int, n int, visit *[10]bool) int {
	for idx, v := range arr {
		visit[v] = true
		if v == n {
			return idx
		}
	}
	return -1
}
