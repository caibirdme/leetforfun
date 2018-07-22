package leet_872

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var ans1, ans2 []int
	getLeafList(root1, &ans1)
	getLeafList(root2, &ans2)
	return arrEq(ans1, ans2)
}

func arrEq(a, b []int) bool {
	length := len(a)
	if len(b) != length {
		return false
	}
	for i := 0; i < length; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func getLeafList(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	if isLeaf(root) {
		*ans = append(*ans, root.Val)
		return
	}
	getLeafList(root.Left, ans)
	getLeafList(root.Right, ans)
}

func isLeaf(root *TreeNode) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	return false
}
