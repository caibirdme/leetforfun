package leet_863

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type item struct {
	val int
	dep int
}

type queue struct {
	buf        []item
	head, tail int
}

func (q *queue) Push(val, dep int) {
	q.buf = append(q.buf, item{val, dep})
	q.tail++
}

func (q *queue) Pop() item {
	t := q.buf[q.head]
	q.head++
	return t
}

func (q *queue) Empty() bool {
	return q.head > q.tail
}

func newQ() *queue {
	return &queue{tail: -1}
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if K == 0 {
		return []int{target.Val}
	}
	var graph [501][501]int
	buildGraph(root, &graph)
	visit := make(map[int]struct{})
	var ans []int
	q := newQ()
	q.Push(target.Val, 0)
	visit[target.Val] = struct{}{}
	for !q.Empty() {
		cur := q.Pop()
		if cur.dep >= K {
			continue
		}
		for i := 0; i <= 500; i++ {
			if graph[cur.val][i] > 0 {
				if _, ok := visit[i]; !ok {
					if cur.dep+1 == K {
						ans = append(ans, i)
						continue
					}
					visit[i] = struct{}{}
					q.Push(i, cur.dep+1)
				}
			}
		}
	}
	return ans
}

func buildGraph(root *TreeNode, graph *[501][501]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		graph[root.Val][root.Left.Val] = 1
		graph[root.Left.Val][root.Val] = 1
		buildGraph(root.Left, graph)
	}
	if root.Right != nil {
		graph[root.Val][root.Right.Val] = 1
		graph[root.Right.Val][root.Val] = 1
		buildGraph(root.Right, graph)
	}
}
