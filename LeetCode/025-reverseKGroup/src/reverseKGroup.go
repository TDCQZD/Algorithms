package src

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	pre := dummy
	end := dummy

	for nil != end.Next {
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}
		if nil == end {
			break
		}

		start := pre.Next
		next := end.Next

		end.Next = nil
		pre.Next = reverse(start)
		start.Next = next

		pre = start
		end = pre
	}
	return dummy.Next
}
func reverse(head *ListNode) *ListNode {
	var pre = &ListNode{}
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

func reverseKGroupByRecursion(head *ListNode, k int) *ListNode {
	if nil == head {
		return head
	}
	swap := head
	for i := 0; i < k; i++ {
		if nil == swap.Next {
			return head
		} else {
			swap = head.Next
		}
	}

	end := swap.Next
	swap.Next = head.Next
	head.Next.Next = head
	head.Next = reverseKGroupByRecursion(end, k)
	head = end

	return swap
}
