package src

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func swapPairsByRecursion(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	next := head.Next
	head.Next = swapPairsByRecursion(next.Next)
	next.Next = head
	return next
}

// 非递归
func swapPairs(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	prev := &ListNode{}
	prev.Next = head
	tmp := prev
	for nil != tmp.Next && nil != tmp.Next.Next {
		start := tmp.Next
		end := tmp.Next.Next
		tmp.Next = end
		start.Next = end.Next
		end.Next = start
		tmp = start
	}
	return prev.Next
}
