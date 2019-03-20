package src

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

// func New(values ...interface{}) *ListNode {
// 	list := &ListNode{}
// 	head := list
// 	if len(values) > 0 {
// 		head.Val = values[0]
// 		for _, value := range values[1:] {
// 			newNode := &ListNode{Val: value}
// 			head.Next = newNode
// 			list = newNode
// 		}
// 	}
// 	return list
// }

func deleteDuplicates(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	prePtr := head
	currentPtr := head.Next
	lastValue := prePtr.Val

	for nil != currentPtr {
		if currentPtr.Val != lastValue {
			lastValue = currentPtr.Val
			prePtr = currentPtr
			currentPtr = currentPtr.Next
		} else {
			prePtr.Next = currentPtr.Next
			currentPtr = currentPtr.Next
		}
	}
	return head
}
