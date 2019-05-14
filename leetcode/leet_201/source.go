package leet_201

func rangeBitwiseAnd(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == n {
		return m
	}
	if m == 1 {
		return 0
	}
	maxi32 := ^uint(0)
	um := uint(m)
	un := uint(n)
	for maxi32&um != maxi32&un {
		maxi32 <<= 1
	}
	return int(maxi32 & um)
}
