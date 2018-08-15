package leet_99

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	var order []*TreeNode
	midTranvel(root, &order)
	a, b := -1, -1
	for i := 0; i < len(order)-1; i++ {
		if order[i].Val > order[i+1].Val {
			a = i
			break
		}
	}
	for i := len(order) - 1; i > a; i-- {
		if order[i].Val < order[i-1].Val {
			b = i
			break
		}
	}
	order[a].Val, order[b].Val = order[b].Val, order[a].Val
}

func midTranvel(root *TreeNode, ans *[]*TreeNode) {
	if root == nil {
		return
	}
	midTranvel(root.Left, ans)
	*ans = append(*ans, root)
	midTranvel(root.Right, ans)
}
