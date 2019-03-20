package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	node1 := &ListNode{Val: 1}
	node2 := &ListNode{Val: 1}
	node3 := &ListNode{Val: 2}
	node4 := &ListNode{Val: 3}
	node5 := &ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	res := deleteDuplicates(node1)
	fmt.Println(res,res.Next,res.Next.Next)
}
