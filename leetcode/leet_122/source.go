package leet_122

func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	ans := 0
	buy := prices[0]
	sell := 0
	buyPhase := true
	var j int
	for i := 1; i < length; i++ {
		if buyPhase {
			for j = i; j < length && buy >= prices[j]; j++ {
				buy = prices[j]
			}
			i = j
			if i < length {
				sell = prices[i]
				buyPhase = false
			}
			i--
		} else {
			for j = i; j < length && sell <= prices[j]; j++ {
				sell = prices[j]
			}
			i = j
			ans += sell - buy
			if i < length {
				buy = prices[i]
				buyPhase = true
			}
			i--
		}
	}
	return ans
}
