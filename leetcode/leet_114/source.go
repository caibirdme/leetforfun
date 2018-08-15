package leet_114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	rotate(root)
}

func rotate(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// leaf
	if root.Left == nil {
		if root.Right == nil {
			return root
		}
		return rotate(root.Right)
	}
	left := root.Left
	llast := rotate(left)
	right := root.Right
	rlast := rotate(right)
	if rlast == nil {
		rlast = llast
	}
	llast.Right = right
	root.Right = left
	root.Left = nil
	return rlast
}
