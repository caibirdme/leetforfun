package leet_546

func removeBoxes(boxes []int) int {
	length := len(boxes)
	if length <= 1 {
		return length
	}
	pre := make([]int, length)
	index := make([]int, 101)
	for i := 0; i <= 100; i++ {
		index[i] = -1
	}
	for idx, b := range boxes {
		pre[idx] = index[b]
		index[b] = idx
	}
	var mem [101][101][101]int
	var mem2 [101][101]int
	return findMax(0, length-1, &mem2, &mem, pre)
}

func findMax(l, r int, mem2 *[101][101]int, mem3 *[101][101][101]int, pre []int) int {
	if l == r {
		return 1
	}
	if l > r {
		return 0
	}
	if t := mem2[l][r]; t > 0 {
		return t
	}
	ans := findMax(l, r-1, mem2, mem3, pre) + 1
	rr := r
	for rr >= l && pre[rr] != -1 && pre[rr] >= l {
		mk := findMaxK(l, pre[rr], pre)
		for k := 1; k <= mk; k++ {
			t := search(l, pre[rr], k, mem2, mem3, pre) + 2*k + 1 + findMax(pre[rr]+1, r-1, mem2, mem3, pre)
			ans = max(ans, t)
		}
		rr = pre[rr]
	}
	mem2[l][r] = ans
	return ans
}

// boxes[l..r]之间有多少个和boxes[r]相等的数
func findMaxK(l, r int, pre []int) int {
	if l == r {
		return 1
	}
	count := 1
	for r >= 0 && pre[r] >= l {
		count++
		r = pre[r]
	}
	return count
}

func search(l, r, k int, mem2 *[101][101]int, mem3 *[101][101][101]int, pre []int) int {
	if l == r {
		return 1
	}
	if l > r {
		return 0
	}
	if t := mem3[l][r][k]; t > 0 {
		return t
	}
	if k == 1 {
		t := findMax(l, r-1, mem2, mem3, pre) + 1
		mem3[l][r][k] = t
		return t
	}
	ans := 0
	rr := r
	for pre[rr] != -1 && pre[rr] >= l {
		mk := findMaxK(l, pre[rr], pre)
		if mk >= k-1 {
			t := search(l, pre[rr], k-1, mem2, mem3, pre) + 2*k - 1 + findMax(pre[rr]+1, r-1, mem2, mem3, pre)
			ans = max(ans, t)
		} else {
			break
		}
		rr = pre[rr]
	}
	mem3[l][r][k] = ans
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func removeBoxes2(boxes []int) int {
	var dp [101][101][101]int
	return dfs(boxes, 0, len(boxes)-1, 0, &dp)
}

func dfs(boxes []int, l, r, k int, dp *[101][101][101]int) int {
	if l > r {
		return 0
	}
	if t := dp[l][r][k]; t > 0 {
		return t
	}
	ans := (k+1)*(k+1) + dfs(boxes, l+1, r, 0, dp)
	for m := l + 1; m <= r; m++ {
		if boxes[l] == boxes[m] {
			ans = max(ans, dfs(boxes, l+1, m-1, 0, dp)+dfs(boxes, m, r, k+1, dp))
		}
	}
	dp[l][r][k] = ans
	return ans
}
