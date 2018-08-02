package leet_877

func stoneGame(piles []int) bool {
	length := len(piles)
	sum := make([]int, length+1)
	for i := 1; i <= length; i++ {
		sum[i] = sum[i-1] + piles[i-1]
	}
	mem := make([][]int, length)
	for i := 0; i < length; i++ {
		mem[i] = make([]int, length)
	}
	score := game(0, length-1, piles, sum, mem)
	if score > sum[length]>>1 {
		return true
	}
	return false
}

func game(i, j int, piles, sum []int, mem [][]int) int {
	if i == j {
		return piles[i]
	}
	if mem[i][j] != 0 {
		return mem[i][j]
	}
	//取第一堆
	all := sum[j+1] - sum[i+1]
	t1 := piles[i] + all - game(i+1, j, piles, sum, mem)
	//取最后一堆
	all = sum[j] - sum[i]
	t2 := piles[j] + all - game(i, j-1, piles, sum, mem)
	mem[i][j] = max(t1, t2)
	return mem[i][j]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
