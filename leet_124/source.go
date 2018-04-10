package leet_124

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := root.Val
	dfs(root, &ans)
	return ans
}

func dfs(root *TreeNode, ans *int) int {
	var l, r int
	v := root.Val
	if root.Left != nil {
		l = dfs(root.Left, ans)
	} else {
		l = -abs(v)
	}
	if root.Right != nil {
		r = dfs(root.Right, ans)
	} else {
		r = -abs(v)
	}
	t := maxMulti(l, r, 0) + v
	*ans = maxMulti(*ans, l, r, l+r+v, t)
	return t
}

func maxMulti(arr ...int) int {
	m := arr[0]
	length := len(arr)
	for i := 1; i < length; i++ {
		m = max(m, arr[i])
	}
	return m
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
