package leet_879

func profitableSchemes(G int, P int, group []int, profit []int) int {
	mod := 1000000007
	length := len(group)
	if length == 0 {
		return 0
	}
	maxP := 0
	for i := 0; i < length; i++ {
		maxP += profit[i]
	}
	f := make([][]int, G+1)
	for i := 0; i <= G; i++ {
		f[i] = make([]int, maxP+1)
	}
	f[0][0] = 1
	for i := 0; i < length; i++ {
		for j := G; j >= group[i]; j-- {
			for k := maxP; k >= profit[i]; k-- {
				f[j][k] += f[j-group[i]][k-profit[i]]
				if f[j][k] >= mod {
					f[j][k] -= mod
				}
			}
		}
	}
	ans := 0
	for i := 1; i <= G; i++ {
		for j := P; j <= maxP; j++ {
			ans += f[i][j]
			if ans >= mod {
				ans -= mod
			}
		}
	}
	return ans
}
