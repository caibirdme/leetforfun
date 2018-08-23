package leet_889

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(pre []int, post []int) *TreeNode {
	length := len(pre)
	if length == 0 {
		return nil
	}
	node := &TreeNode{
		Val: pre[0],
	}
	if length == 1 {
		return node
	}
	pre = pre[1:]
	post = post[:length-1]
	leftVal := pre[0]
	idx := findVal(post, leftVal)
	node.Left = constructFromPrePost(pre[:idx+1], post[:idx+1])
	node.Right = constructFromPrePost(pre[idx+1:], post[idx+1:])
	return node
}

func findVal(post []int, val int) int {
	for idx, v := range post {
		if v == val {
			return idx
		}
	}
	return -1
}
