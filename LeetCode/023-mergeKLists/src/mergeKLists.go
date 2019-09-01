package src

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if nil == lists || 0 == len(lists) {
		return nil
	}
	if 1 == len(lists) {
		return lists[0]
	}
	l := len(lists)
	i := 1
	for i < l {
		for j := 0; j < l-i; j += i * 2 {
			lists[j] = mergeTwoLists(lists[j], lists[j+i])
		}

		i *= 2
	}

	return lists[0]
}
func mergeKLists2(lists []*ListNode) *ListNode {
	if nil == lists || 0 == len(lists) {
		return nil
	}

	if 1 == len(lists) {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	list := mergeTwoLists(lists[0], lists[1])
	for i := 2; i < len(lists); i++ {
		list = mergeTwoLists(list, lists[i])
	}
	return list
}
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	if nil == l1 {
		return l2
	}
	if nil == l2 {
		return l1
	}

	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}
