package leet_903

func numPermsDISequence(S string) int {
	length := len(S)
	if length <= 1 {
		return length
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, length)
	}
	return dfs(0, length-1, []byte(S), dp)
}

const mod = int(1e9 + 7)

func dfs(l, r int, str []byte, dp [][]int) int {
	if l >= r {
		return 1
	}
	if dp[l][r] > 0 {
		return dp[l][r]
	}
	if str[l] == 'I' {
		dp[l][r] = dfs(l+1, r, str, dp)
		dp[l][r] %= mod
	}
	if str[r] == 'D' {
		dp[l][r] += dfs(l, r-1, str, dp)
		dp[l][r] %= mod
	}
	n := r - l + 2
	for i := l + 1; i <= r; i++ {
		if str[i-1] == 'D' && str[i] == 'I' {
			dp[l][r] += calc.C(n-1, i-l) * dfs(l, i-2, str, dp) % mod * dfs(i+1, r, str, dp) % mod
			dp[l][r] %= mod
		}
	}
	return dp[l][r]
}

var calc = memory{f: []int{1, 1}, facinv: make(map[int]int)}

type memory struct {
	f      []int
	facinv map[int]int
}

func (b *memory) C(n, m int) int {
	if m == 0 || m == n {
		return 1
	}
	if m == 1 || m == n-1 {
		return n
	}
	t := b.fac(n) * b.inv(b.fac(m)) % mod * b.inv(b.fac(n-m)) % mod
	return t
}

func (b *memory) inv(n int) int {
	if v, ok := b.facinv[n]; ok {
		return v
	}
	t := pow(n, mod-2, mod)
	b.facinv[n] = t
	return t
}

func pow(a, b, m int) int {
	if b < 0 {
		return 0
	}
	result := 1
	a %= m
	for b > 0 {
		if b&1 == 1 {
			result = (result * a) % m
		}
		b >>= 1
		a = (a * a) % m
	}
	return result
}

func (b *memory) fac(n int) int {
	last := len(b.f) - 1
	if n <= last {
		return b.f[n]
	}
	for i := last + 1; i <= n; i++ {
		b.f = append(b.f, (b.f[i-1]*i)%mod)
	}
	return b.f[n]
}
