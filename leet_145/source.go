package leet_145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	s := newStack()
	s.Push(root)
	visited := make(map[*TreeNode]struct{})
	var ans []int
	for !s.Empty() {
		node := s.Pop()
		if nil == node {
			continue
		}
		if isLeaf(node) {
			ans = append(ans, node.Val)
			continue
		}
		if _, ok := visited[node]; ok {
			ans = append(ans, node.Val)
			continue
		}
		visited[node] = struct{}{}
		s.Push(node)
		s.Push(node.Right)
		s.Push(node.Left)
	}
	return ans
}

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

type stack struct {
	buf []*TreeNode
	sp  int
}

func newStack() *stack {
	return &stack{sp: -1}
}

func (s *stack) Empty() bool {
	return s.sp < 0
}

func (s *stack) Push(node *TreeNode) {
	s.sp++
	if len(s.buf) > s.sp {
		s.buf[s.sp] = node
		return
	}
	s.buf = append(s.buf, node)
}

func (s *stack) Top() *TreeNode {
	return s.buf[s.sp]
}

func (s *stack) Pop() *TreeNode {
	if s.sp >= 0 {
		t := s.buf[s.sp]
		s.sp--
		return t
	}
	return nil
}
