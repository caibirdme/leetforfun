package leet_188

func maxProfit(k int, prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	if k == 0 {
		return 0
	}
	for i := length - 1; i >= 1; i-- {
		prices[i] = prices[i] - prices[i-1]
	}
	prices[0] = 0
	if length/2 <= k {
		return maxProfitUnlimited(prices)
	}
	/*
		f := make([][2]int, length)
		g := make([][2]int, length)
		cur := 0
		for i := 1; i <= k; i++ {
			for i := 1; i < length; i++ {
				f[i][cur] = max(f[i-1][cur]+prices[i], g[i-1][cur^1])
				g[i][cur] = max(f[i][cur], g[i-1][cur])
			}
			cur ^= 1
		}

		return max(g[length-1][0], g[length-1][1])
	*/
	f := make([][2][2]int, length)
	cur := 0
	for j := 1; j <= k; j++ {
		for i := 1; i < length; i++ {
			f[i][cur][1] = max(f[i-1][cur][1], f[i-1][cur^1][0]) + prices[i]
			f[i][cur][0] = max(f[i-1][cur][0], f[i-1][cur][1])
		}
		cur ^= 1
	}
	last := length - 1
	return max(f[last][0][0], f[last][0][1], f[last][1][0], f[last][1][1])
}

func max(arr ...int) int {
	if len(arr) == 1 {
		return arr[0]
	}
	t := max(arr[1:]...)
	if arr[0] > t {
		return arr[0]
	}
	return t
}

func maxProfitUnlimited(diff []int) int {
	var ans int
	length := len(diff)
	for i := 1; i < length; i++ {
		if diff[i] > 0 {
			ans += diff[i]
		}
	}
	return ans
}
