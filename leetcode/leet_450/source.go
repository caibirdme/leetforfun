package leet_450

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	doDelete(&root, key)
	return root
}

func doDelete(node **TreeNode, key int) {
	if node == nil {
		return
	}
	cur := *node
	if cur == nil {
		return
	}
	if key == cur.Val {
		if isLeaf(cur) {
			*node = nil
			return
		}
		if sc := getSingleChild(cur); sc != nil {
			*node = sc
			return
		}
		r := cur.Right
		if r.Left == nil {
			r.Left = cur.Left
			*node = r
			return
		}
		leftMost := findLeftMost(r)
		(*node).Val = leftMost.Val
	} else if key < cur.Val {
		doDelete(&cur.Left, key)
	} else {
		doDelete(&cur.Right, key)
	}
}

func getSingleChild(node *TreeNode) *TreeNode {
	if node.Left != nil && node.Right == nil {
		return node.Left
	}
	if node.Right != nil && node.Left == nil {
		return node.Right
	}
	return nil
}

// 找到最左边的子树
func findLeftMost(node *TreeNode) *TreeNode {
	parent := node
	cur := parent.Left
	for cur.Left != nil {
		parent = cur
		cur = cur.Left
	}
	parent.Left = cur.Right
	cur.Right = nil
	return cur
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
