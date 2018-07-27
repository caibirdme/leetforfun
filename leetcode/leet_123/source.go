package leet_123

func maxProfit(prices []int) int {
	ans := singleTransaction(prices)
	if ans == 0 {
		return 0
	}
	for i := 2; i < len(prices)-1; i++ {
		// 如果要进行两次交易，那么第一次交易必然是高位套现
		// 假设第一次在prices[i]卖出，如果prices[i+1]比prices[i]价格更高
		// 那显然在prices[i]卖出就不是最优的选择
		// 因此第一次卖一定是在一个高点，即prices[i]>prices[i+1]
		// 每次枚举这个高点，把prices分成两部分，分别进行处理
		if prices[i] < prices[i-1] {
			ans = max(ans, singleTransaction(prices[:i])+singleTransaction(prices[i:]))
		}
	}
	return ans
}

func singleTransaction(prices []int) int {
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
