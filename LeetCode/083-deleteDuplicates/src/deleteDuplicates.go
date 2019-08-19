package src

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

// 直接法
func deleteDuplicates3(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	curr := head
	for nil != curr && nil != curr.Next {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head
}

// 双指针
func deleteDuplicates2(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	curr := head
	fast := head.Next
	for nil != fast {
		if curr.Val == fast.Val {
			curr.Next = fast.Next
			fast = curr.Next
		} else {
			curr = fast
			fast = curr.Next
		}
	}
	return head
}

// 递归
func deleteDuplicates(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		head = head.Next
	}
	return head

}
