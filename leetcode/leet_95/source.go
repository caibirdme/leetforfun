package leet_95

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return gen(1, n)
}

func gen(l, r int) []*TreeNode {
	if l == r {
		return []*TreeNode{&TreeNode{Val: l}}
	}
	if l > r {
		return nil
	}
	var ans []*TreeNode
	for i := l; i <= r; i++ {
		left := i - 1
		right := i + 1
		leftSubTree := gen(l, left)
		rightSubTree := gen(right, r)
		if len(leftSubTree) == 0 {
			for k := 0; k < len(rightSubTree); k++ {
				node := new(TreeNode)
				node.Right = rightSubTree[k]
				node.Val = i
				ans = append(ans, node)
			}
		} else if len(rightSubTree) == 0 {
			for k := 0; k < len(leftSubTree); k++ {
				node := new(TreeNode)
				node.Val = i
				node.Left = leftSubTree[k]
				ans = append(ans, node)
			}
		} else {
			for j := 0; j < len(leftSubTree); j++ {
				for k := 0; k < len(rightSubTree); k++ {
					node := new(TreeNode)
					node.Val = i
					node.Left = leftSubTree[j]
					node.Right = rightSubTree[k]
					ans = append(ans, node)
				}
			}
		}
	}
	return ans
}
