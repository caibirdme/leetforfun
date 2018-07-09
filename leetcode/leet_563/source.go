package leet_563

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	var tilt int
	visit(root, &tilt)
	return tilt
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func visit(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}
	leftSum := visit(root.Left, tilt)
	rightSum := visit(root.Right, tilt)
	*tilt += abs(leftSum - rightSum)
	return leftSum + rightSum + root.Val
}
