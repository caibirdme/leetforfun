package leet_105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(&preorder, inorder)
}

func build(preorder *[]int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	node := new(TreeNode)
	node.Val = (*preorder)[0]
	var pos int
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == node.Val {
			pos = i
			break
		}
	}
	*preorder = (*preorder)[1:]
	node.Left = build(preorder, inorder[:pos])
	node.Right = build(preorder, inorder[pos+1:])
	return node
}
