package utils

import (
	"fmt"
)

/*链表实现队列*/

/*单向环行链表*/

type CatNode struct{
	id int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode)  {
	
	//判断是否是第一只
	 if head.next == nil{
		 head.id = newCatNode.id
		 head.name = newCatNode.name
		 head.next = head //形成环
		 return
	 }else{
		 temp := head
		for {
			
			if temp.next == head {
				break
			}
			temp = temp.next
		}
		temp.next = newCatNode
		newCatNode.next = head
		
	 }
}

//删除
func DelCatNode(head *CatNode, id int) *CatNode{
	temp := head
	helper := head
	flag := true 
	//空链表
	if temp.next == nil {
		return head
	}
	//只有一个节点
	if temp.next == head {
		if temp.id == id {
		temp.next = nil	
		}
		return head
	}
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}
	// 两个以上节点
	for{
		if temp.next == head {
			break
		}
		if temp.id == id {
			if temp == head {
				head =head.next
			}
			helper.next = temp.next
			flag = false
			break
		}
		temp =temp.next
		helper = helper.next		
	}

	if flag{
		if temp.id == id {
			helper.next = temp.next
		}
	
	}else{
		fmt.Println("ID不存在")
	}
	return head
}


func ShowLinkedListC(head *CatNode)  {
	temp := head
	if temp.next == nil {
		return
	}
	for{
		fmt.Print("==>",temp)
		
		if temp.next == head {
			break
		}
		temp =temp.next
	}
}

func CircleLinkedList()  {

	head := &CatNode{}
	cat1 := &CatNode{
		id : 1,
		name : "小花",
	}
	cat2 := &CatNode{
		id : 2,
		name : "小花",
	}
	cat3 := &CatNode{
		id : 3,
		name : "小花",
	}
	cat4 := &CatNode{
		id : 4,
		name : "小花",
	}
	InsertCatNode(head,cat1)
	InsertCatNode(head,cat2)
	InsertCatNode(head,cat3)
	InsertCatNode(head,cat4)
	ShowLinkedListC(head)
	head = DelCatNode(head,1)
	ShowLinkedListC(head)
}