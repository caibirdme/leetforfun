package testsuit

type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := new(ListNode)
	head.Val = arr[0]
	last := head
	for i := 1; i < len(arr); i++ {
		current := new(ListNode)
		current.Val = arr[i]
		last.Next = current
		last = current
	}
	return head
}

func CmpList(a, b *ListNode) bool {
	currentA, currentB := a, b
	for {
		if currentA == nil && currentB != nil {
			return false
		}
		if currentB == nil && currentA != nil {
			return false
		}
		if currentA == nil && currentB == nil {
			return true
		}
		if currentA.Val != currentB.Val {
			return false
		}
		currentA = currentA.Next
		currentB = currentB.Next
	}
}
