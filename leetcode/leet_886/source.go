package leet_886

func reachableNodes(edges [][]int, M int, N int) int {
	if len(edges) == 0 || N == 1 {
		return 1
	}
	connect := make(map[int]map[int]int)
	updateMap := make(map[int]map[int]int)
	for _, edge := range edges {
		x, y, v := edge[0], edge[1], edge[2]
		if _, ok := connect[x]; !ok {
			connect[x] = make(map[int]int)
			updateMap[x] = make(map[int]int)
		}
		if _, ok := connect[y]; !ok {
			connect[y] = make(map[int]int)
			updateMap[y] = make(map[int]int)
		}
		v++
		connect[x][y] = v
		connect[y][x] = v
		updateMap[x][y] = v
		updateMap[y][x] = v
	}

	dist := findShortestPath(connect, N)
	q := newQ()
	q.Push(elem{distance: 0, node: 0})
	calculated := map[int]struct{}{
		0: struct{}{},
	}
	ans := 1
	for !q.Empty() {
		e := q.Pop().(elem)
		edgeList, ok := connect[e.node]
		if !ok {
			continue
		}
		for node, distance := range edgeList {
			// 是最短路
			if _, ok := calculated[node]; !ok && e.distance+distance == dist[node] {
				if dist[node] >= M {
					d := M - e.distance
					updateMap[e.node][node] -= d
					updateMap[node][e.node] -= d
					ans += d
				} else {
					ans += distance
					delete(connect[e.node], node)
					delete(connect[node], e.node)
					q.Push(elem{distance: e.distance + distance, node: node})
					calculated[node] = struct{}{}
				}
			} else {
				d := min(updateMap[e.node][node]-1, M-e.distance)
				updateMap[e.node][node] -= d
				updateMap[node][e.node] -= d
				ans += d
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type elem struct {
	distance int
	node     int
}

func findShortestPath(connect map[int]map[int]int, N int) []int {
	dist := make([]int, N)
	for i := 1; i < N; i++ {
		dist[i] = _maxInt
	}
	q := newQ()
	q.Push(0)
	for !q.Empty() {
		node := q.Pop().(int)
		edgeList, ok := connect[node]
		if !ok {
			continue
		}
		for nextNode, distance := range edgeList {
			if dist[node]+distance < dist[nextNode] {
				dist[nextNode] = dist[node] + distance
				q.Push(nextNode)
			}
		}
	}
	return dist
}

type queue struct {
	head int
	list []interface{}
}

func (q *queue) Push(v interface{}) {
	q.list = append(q.list, v)
}

func (q *queue) Pop() interface{} {
	q.head++
	v := q.list[q.head]
	return v
}

func (q *queue) Empty() bool {
	return q.head >= len(q.list)-1
}

func newQ() *queue {
	return &queue{head: -1}
}

const _maxInt = 0xffffffff
