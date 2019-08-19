package src

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
func swapPairsByRecursion(head *ListNode) *ListNode {
	//终止条件：链表只剩一个节点或者没节点了，没得交换了。返回的是已经处理好的链表
	if nil == head || nil == head.Next {
		return head
	}
	//一共三个节点:head, next, swapPairs(next.next)
	//下面的任务便是交换这3个节点中的前两个节点
	next := head.Next
	head.Next = swapPairsByRecursion(next.Next)
	next.Next = head
	//根据第二步：返回给上一级的是当前已经完成交换后，即处理好了的链表部分
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
