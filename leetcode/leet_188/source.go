package leet_188

func maxProfit(k int, prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	if k == 0 {
		return 0
	}
	var bf, bg []int
	if prices[0] <= prices[1] {
		bg = append(bg, 1)
	}
	//寻找波峰、波谷
	for i := 2; i < length; i++ {
		idx := i - 1
		if prices[idx-1] <= prices[idx] && prices[idx] > prices[idx+1] {
			//波峰
			bf = append(bf, i)
		} else if prices[idx-1] >= prices[idx] && prices[idx] < prices[idx+1] {
			//波谷
			bg = append(bg, i)
		}
	}
	if prices[length-1] >= prices[length-2] {
		bf = append(bf, length)
	}
	bglen := len(bg)
	bflen := len(bf)
	if bflen == 0 || bglen == 0 {
		return 0
	}
	for bf[0] < bg[0] {
		bf = bf[1:]
	}
	bflen = len(bf)

	if k > bflen {
		k = bflen
	}
	if k > bglen {
		k = bglen
	}
	length = bflen + bglen
	newPrice := make([]int, length)
	{
		i, j := 0, 0
		for i+j < length {
			if j >= bglen {
				newPrice[i+j] = prices[bf[i]-1]
				i++
			} else if i >= bflen {
				newPrice[i+j] = prices[bg[j]-1]
				j++
			} else if bf[i] < bg[j] {
				newPrice[i+j] = prices[bf[i]-1]
				i++
			} else {
				newPrice[i+j] = prices[bg[j]-1]
				j++
			}
		}
	}
	f := make([][]int, length+1)
	for i := 0; i <= length; i++ {
		f[i] = make([]int, k+1)
	}
	for i := 2; i <= length; i++ {
		idx := i - 1
		mt := i >> 1
		for j := 1; j <= k && j <= mt; j++ {
			f[i][j] = f[i-1][j]
		}
		if newPrice[idx] <= newPrice[idx-1] {
			continue
		}
		sell := newPrice[idx]
		buy := newPrice[idx-1]
		profit := sell - buy
		for t := 1; i <= k && t <= mt; t++ {

		}
	}
	ans := 0
	for i := 1; i <= k; i++ {
		ans = max(ans, f[length][i])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
