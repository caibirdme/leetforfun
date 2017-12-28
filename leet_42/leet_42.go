package leet_42

// segment tree
func trap(height []int) int {
	length := len(height)
	if length < 3 {
		return 0
	}
	sum := 0
	tree := newSegmentTree(height)
	for i := 1; i < length-1; i++ {
		l := tree.search(0, i-1)
		r := tree.search(i+1, length-1)
		h := min(l, r)
		sum += max(0, h-height[i])
	}
	return sum
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type segmentTree struct {
	l, r   int
	data   int
	lc, rc *segmentTree
}

func (t *segmentTree) search(l, r int) int {
	if t.l == l && t.r == r {
		return t.data
	}
	mid := (t.l + t.r) >> 1
	if r <= mid {
		return t.lc.search(l, r)
	}
	if l >= mid+1 {
		return t.rc.search(l, r)
	}
	return max(t.lc.search(l, mid), t.rc.search(mid+1, r))
}

func newSegmentTree(arr []int) *segmentTree {
	return buildTree(arr, 0, len(arr)-1)
}

func buildTree(arr []int, l, r int) *segmentTree {
	node := &segmentTree{
		l: l,
		r: r,
	}
	if l == r {
		node.data = arr[l]
		return node
	}
	mid := (l + r) >> 1
	node.lc = buildTree(arr, l, mid)
	node.rc = buildTree(arr, mid+1, r)
	node.data = max(node.lc.data, node.rc.data)
	return node
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
