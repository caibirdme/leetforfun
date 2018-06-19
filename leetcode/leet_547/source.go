package leet_547

func findCircleNum(M [][]int) int {
	length := len(M)
	parent := make([]int, length)
	for i := 0; i < length; i++ {
		parent[i] = i
	}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if M[i][j] == 1 {
				mergeParent(i, j, parent)
			}
		}
	}
	group := make(map[int]struct{})
	count := 0
	for i := 0; i < length; i++ {
		t := findParent(i, parent)
		if _, ok := group[t]; !ok {
			group[t] = struct{}{}
			count++
		}
	}
	return count
}

func findParent(i int, parent []int) int {
	if parent[i] == i {
		return i
	}
	parent[i] = findParent(parent[i], parent)
	return parent[i]
}

func mergeParent(i, j int, parent []int) {
	iParent := findParent(i, parent)
	jParent := findParent(j, parent)
	parent[jParent] = iParent
}
