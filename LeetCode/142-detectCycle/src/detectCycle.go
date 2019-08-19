package src

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//  哈希表
func detectCycle(head *ListNode) *ListNode {
	// if nil == head || nil == head.Next {
	// 	return head
	// }
	hash := make(map[*ListNode]int)
	for head != nil {
		if _, ok := hash[head]; ok {
			return head
		}
		hash[head] = head.Val
		head = head.Next
	}
	return nil
}

//
func detectCycle2(head *ListNode) *ListNode {
	if nil == head || nil == head.Next {
		return nil
	}
	intersect := getIntersect(head)
	if nil == intersect {
		return nil
	}

	ptr1 := head
	ptr2 := intersect
	for ptr1 != ptr2 {
		ptr1 = ptr1.Next
		ptr2 = ptr2.Next
	}

	return ptr1

}
func getIntersect(head *ListNode) *ListNode {
	tortoise := head
	hare := head
	for hare != nil && hare.Next != nil {
		tortoise = tortoise.Next
		hare = hare.Next.Next // 慢指针，每次走一步
		if tortoise == hare { // 快慢指针相遇，表示有环
			return tortoise
		}

	}
	return nil
}
