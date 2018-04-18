package leet_233

var count = [9]int{0, 1}
var pow = [10]int{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}

func init() {
	mul := 1
	for i := 2; i <= 8; i++ {
		mul *= 10
		count[i] = mul + 10*count[i-1]
	}
}

func countDigitOne(n int) int {
	if n <= 9 {
		if n >= 1 {
			return 1
		}
		return 0
	}
	length := getLength(n)
	base := pow[length-1]
	firstNum := n / base
	rest := n % base
	if firstNum == 1 {
		// 1
		return rest + 1 + countDigitOne(rest) + count[length-1]
	}
	// >=2
	return base + firstNum*count[length-1] + countDigitOne(rest)
}

// find the maximal 10^t which is less equal than n
// if n = 456, the p should be 100
// if n = 8, the p should be 1
// if n = 10, the p should be 10
func getBase(n int) int {
	t := 1
	p := 1
	for p*10 <= n {
		t++
		p *= 10
	}
	return p
}

func getLength(n int) int {
	l := 0
	for n > 0 {
		l++
		n /= 10
	}
	return l
}
