package leet_743

const unreachable = -1

func networkDelayTime(times [][]int, N int, K int) int {
	dist := make([]int, N+1)
	graph := make([][]int, N+1)
	for i := 1; i <= N; i++ {
		dist[i] = unreachable
		graph[i] = make([]int, N+1)
		for j := 1; j <= N; j++ {
			graph[i][j] = unreachable
		}
	}
	for _, uvw := range times {
		u, v, w := uvw[0], uvw[1], uvw[2]
		if graph[u][v] == unreachable {
			graph[u][v] = w
		} else {
			graph[u][v] = min(graph[u][v], w)
		}
	}
	dist[K] = 0
	q := new(queue)
	q.tail = -1
	q.Push(K)
	for !q.Empty() {
		node := q.Pop()
		for i := 1; i <= N; i++ {
			if node == i || graph[node][i] == unreachable {
				continue
			}
			if dist[i] == unreachable {
				dist[i] = dist[node] + graph[node][i]
				q.Push(i)
			} else {
				cost := dist[node] + graph[node][i]
				if cost < dist[i] {
					dist[i] = cost
					q.Push(i)
				}
			}
		}
	}
	ans := -1
	for i := 1; i <= N; i++ {
		if dist[i] == unreachable {
			return -1
		}
		ans = max(ans, dist[i])
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

type queue struct {
	buf        []int
	head, tail int
}

func (q *queue) Push(i int) {
	q.buf = append(q.buf, i)
	q.tail++
}

func (q *queue) Pop() int {
	defer func() {
		q.head++
	}()
	return q.buf[q.head]
}

func (q *queue) Empty() bool {
	return q.head > q.tail
}
