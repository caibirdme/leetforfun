package leet_134

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	n := length * 2
	extendGas := make([]int, n)
	extendCost := make([]int, n)
	sumGas := make([]int, n+1)
	sumCost := make([]int, n+1)
	for i := 0; i < length; i++ {
		t := i + length
		extendGas[i], extendGas[t] = gas[i], gas[i]
		extendCost[i], extendCost[t] = cost[i], cost[i]
	}
	for i := 1; i <= n; i++ {
		sumGas[i] = sumGas[i-1] + extendGas[i-1]
		sumCost[i] = sumCost[i-1] + extendCost[i-1]
	}
	q := &queue{head: -1}
	count := 0
	rest := 0
	start := -1
	for i := 0; i < n; i++ {
		if rest == 0 {
			if extendGas[i] < extendCost[i] {
				continue
			} else {
				rest += extendGas[i]
				start = i
			}
		} else {
			k := -1
			for rest < extendCost[i-1] && !q.Empty() {
				k = q.Pop()
				rest = rest - (sumGas[k-1] - sumGas[start]) + sumCost[k-1] - sumCost[start-1]
				count = i - k
			}
			if rest < extendCost[i-1] {
				count = 0
				start = -1
				rest = 0
				i--
			} else {
				if k != -1 {
					start = k
				}
				rest = rest - extendCost[i-1] + extendGas[i]
				count++
				if count == length {
					return start
				}
			}
		}
	}
	return -1
}

type queue struct {
	buf  []int
	head int
}

func (q *queue) Append(n int) {
	q.buf = append(q.buf, n)
}

func (q *queue) Pop() int {
	q.head++
	return q.buf[q.head]
}

func (q *queue) Empty() bool {
	return q.head+1 >= len(q.buf)
}
