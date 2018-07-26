package leet_121

func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	ans := 0
	buy := prices[0]
	for i := 1; i < length; i++ {
		if prices[i] > buy {
			ans = max(ans, prices[i]-buy)
		} else {
			buy = prices[i]
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
