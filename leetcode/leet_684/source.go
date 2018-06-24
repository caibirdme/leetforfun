package leet_684

func findRedundantConnection(edges [][]int) []int {
	N := len(edges)
	parent := make([]int, N+1)
	for i := 1; i <= N; i++ {
		parent[i] = i
	}
	var ans []int
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		pX, pY := findParent(x, parent), findParent(y, parent)
		if pX == pY {
			if 0 == len(ans) {
				ans = append(ans, x, y)
			} else {
				ans[0], ans[1] = x, y
			}
		} else {
			mergeParent(pX, pY, parent)
		}
	}
	return ans
}

func findParent(i int, parent []int) int {
	if parent[i] == i {
		return i
	}
	parent[i] = findParent(parent[i], parent)
	return parent[i]
}

func mergeParent(i, j int, parent []int) {
	pI, pJ := findParent(i, parent), findParent(j, parent)
	parent[pJ] = pI
}
