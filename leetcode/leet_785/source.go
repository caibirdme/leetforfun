package leet_785

func isBipartite(graph [][]int) bool {
	var color = [2]map[int]struct{}{
		make(map[int]struct{}),
		make(map[int]struct{}),
	}
	q := newQ()
	for i := 0; i < len(graph); i++ {
		if len(graph[i]) == 0 {
			continue
		}
		_, ok1 := color[0][i]
		_, ok2 := color[1][i]
		if ok1 || ok2 {
			continue
		}
		q.Push(node{idx: i, pos: 0})
		for !q.Empty() {
			cur := q.Pop()
			pos := 1 - cur.pos
			for _, v := range graph[cur.idx] {
				if _, ok := color[cur.pos][v]; ok {
					return false
				}
				if _, ok := color[pos][v]; ok {
					continue
				} else {
					color[pos][v] = struct{}{}
					q.Push(node{idx: v, pos: pos})
				}
			}
		}
	}
	return true
}

type node struct {
	idx int
	pos int
}

type queue struct {
	buf        []node
	head, tail int
}

func newQ() *queue {
	return &queue{tail: -1}
}

func (q *queue) Push(n node) {
	q.tail++
	q.buf = append(q.buf, n)
}

func (q *queue) Pop() node {
	t := q.buf[q.head]
	q.head++
	return t
}

func (q *queue) Empty() bool {
	return q.head > q.tail
}
