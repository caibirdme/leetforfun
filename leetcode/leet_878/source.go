package leet_878

func nthMagicalNumber(N int, A int, B int) int {
	const mod = 1000000007
	if A > B {
		A, B = B, A
	}
	if A == B || B%A == 0 {
		return (A * N) % mod
	}
	t1 := A
	t2 := B
	step := lcm(A, B)
	group := step/A + step/B - 1
	if N >= group {
		count := N / group * group
		last := N / group * step
		if count == N {
			return last
		}
		t1 += last
		t2 += last
		if t1 >= mod && t2 >= mod {
			t1 %= mod
			t2 %= mod
		}
		N -= count
	}
	ans := t1
	for i := 0; i < N; i++ {
		if t1 < t2 {
			ans = t1
			t1 += A
		} else if t1 > t2 {
			ans = t2
			t2 += B
		} else {
			ans = t1
			t1 += A
			t2 += B
		}
		if t1 >= mod && t2 >= mod {
			t1 -= mod
			t2 -= mod
		}
	}
	return ans
}

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
