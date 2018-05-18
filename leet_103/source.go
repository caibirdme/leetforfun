package leet_103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	if nil == root {
		return nil
	}
	var ans [][]int
	var queue [2][]*TreeNode
	queue[0] = []*TreeNode{root}
	order := 0
	for {
		length := len(queue[order])
		if 0 == length {
			break
		}
		buff := make([]int, length)
		var idx, delta int
		if order&1 == 1 {
			idx = length - 1
			delta = -1
		} else {
			idx = 0
			delta = 1
		}
		next := 1 - order
		for len(queue[order]) > 0 {
			node := queue[order][0]
			buff[idx] = node.Val
			idx += delta
			if node.Left != nil {
				queue[next] = append(queue[next], node.Left)
			}
			if node.Right != nil {
				queue[next] = append(queue[next], node.Right)
			}
			if len(queue[order]) > 1 {
				queue[order] = queue[order][1:]
			} else {
				queue[order] = queue[order][:0]
			}
		}
		ans = append(ans, append([]int(nil), buff...))
		queue[order] = queue[order][:0]
		order = next
	}
	return ans
}
