package leet_98

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isValid(root.Left, unlimit, root.Val) && isValid(root.Right, root.Val, unlimit)
}

const unlimit = 0x7f7f7f7f

func isValid(cur *TreeNode, min, max int) bool {
	if cur == nil {
		return true
	}
	if min != unlimit && cur.Val <= min {
		return false
	}
	if max != unlimit && cur.Val >= max {
		return false
	}
	if cur.Left != nil {
		if cur.Left.Val >= cur.Val {
			return false
		}
		if !isValid(cur.Left, min, cur.Val) {
			return false
		}
	}
	if cur.Right != nil {
		if cur.Right.Val <= cur.Val {
			return false
		}
		if !isValid(cur.Right, cur.Val, max) {
			return false
		}
	}
	return true
}
