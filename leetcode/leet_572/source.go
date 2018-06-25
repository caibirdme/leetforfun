package leet_572

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil {
		return false
	}

	return visit(s, t) || isSubtree(s.Left, t) || isSubtree(s.Right, t)
}

func visit(s *TreeNode, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	return visit(s.Left, t.Left) && visit(s.Right, t.Right)
}
