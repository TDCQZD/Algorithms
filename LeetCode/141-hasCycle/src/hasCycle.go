package src

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]int) // 开一个哈希表记录该节点是否已经遍历过，值记录节点索引
	for head != nil {
		if _, ok := hash[head]; ok { // 该节点遍历过，形成了环
			return true
		}
		hash[head] = head.Val // 记录该节点已经遍历过
		head = head.Next
	}
	return false
}
func hasCycle2(head *ListNode) bool { // 快慢指针。假如爱有天意，那么快慢指针终会相遇
	if nil == head || nil == head.Next {
		return false
	}
	fastHead := head.Next // 快指针，每次走两步
	for fastHead != nil && head != nil && fastHead.Next != nil {
		if fastHead == head { // 快慢指针相遇，表示有环
			return true
		}
		fastHead = fastHead.Next.Next
		head = head.Next // 慢指针，每次走一步
	}
	return false
}
func hasCycle3(head *ListNode) bool { // 走自己的路让别人无路可走思路
	for head != nil {
		if head.Val == 18464187 { // 这是自己走过的路，说明有环
			return true
		}
		head.Val = 18464187
		head = head.Next
	}
	return false
}
