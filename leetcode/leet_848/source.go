package leet_848

func shiftingLetters(S string, shifts []int) string {
	root := newTree(S)
	for idx, delta := range shifts {
		root.add(0, idx, delta)
	}
	arr := make([]int, len(S))
	root.full(arr)
	bytes := make([]byte, len(S))
	posA := int('a')
	for idx, delta := range arr {
		bytes[idx] = byte((delta-posA)%26) + 'a'
	}
	return string(bytes)
}

type treeNode struct {
	val         int
	delta       int
	left, right int
	lch, rch    *treeNode
}

func (t *treeNode) add(l, r int, delta int) {
	if t.left == l && t.right == r {
		t.delta += delta
		return
	}
	mid := (t.left + t.right) >> 1
	delta += t.delta
	if r <= mid {
		t.lch.add(l, r, delta)
	} else if l >= mid+1 {
		t.rch.add(l, r, delta)
	} else {
		t.lch.add(l, mid, delta)
		t.rch.add(mid+1, r, delta)
	}
	t.delta = 0
}

func (t *treeNode) full(arr []int) {
	if t.left == t.right {
		arr[t.left] = t.val + t.delta
		return
	}
	if t.delta > 0 {
		t.lch.delta += t.delta
		t.rch.delta += t.delta
		t.delta = 0
	}
	t.lch.full(arr)
	t.rch.full(arr)
}

func newTree(s string) *treeNode {
	length := len(s)
	if 0 == length {
		return nil
	}
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[i] = int(s[i])
	}
	root := new(treeNode)
	buildTree(root, 0, length-1, arr)
	return root
}

func buildTree(root *treeNode, l, r int, arr []int) {
	if l == r {
		root.left = l
		root.right = l
		root.val = arr[l]
		return
	}
	root.left = l
	root.right = r
	mid := (l + r) >> 1
	root.lch = new(treeNode)
	root.rch = new(treeNode)
	buildTree(root.lch, l, mid, arr)
	buildTree(root.rch, mid+1, r, arr)
}
