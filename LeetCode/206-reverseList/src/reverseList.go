package src

type ListNode struct {
	Val  int
	Next *ListNode
}

// 迭代实现方式一：借用两个辅助变量 最优解
func reverseListByIteration1(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}

	dummy := head
	tmp := head

	for head != nil && head.Next != nil {
		dummy = head.Next
		head.Next = dummy.Next
		dummy.Next = tmp
		tmp = dummy
	}

	return dummy
}

// 迭代实现方式二：构建新链表
func reverseListByIteration2(head *ListNode) *ListNode {
	var newHead *ListNode
	for nil != head {
		nextNode := head.Next
		head.Next = newHead
		newHead = head
		head = nextNode
	}
	return newHead
}

// 迭代实现方式一：
func reverseListByRecursive1(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return head
	}
	p := reverseListByRecursive1(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

// 迭代实现方式二：构建新链表
func reverseListByRecursive2(head *ListNode) *ListNode {
	return reverseCore(head, nil)
}

func reverseCore(head, newHead *ListNode) *ListNode {
	if nil == head {
		return newHead
	}

	nextNode := head.Next
	head.Next = newHead
	return reverseCore(nextNode, head)
}

// 迭代实现方式二：构建新链表
func reverseListByRecursive3(head *ListNode) *ListNode {

	var prev *ListNode
	curr := head
	for nil != curr {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
