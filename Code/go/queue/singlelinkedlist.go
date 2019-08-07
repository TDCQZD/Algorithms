package utils

import (
	"fmt"
)

/*链表实现队列*/

/*单链表*/

type HeroNode struct{
	//数据本身
	id int
	name string
	nickName string
	//nextNode指向一个和它本身存储指向下一个节点的指针
	nextNode *HeroNode
}

//无序插入
//链表插入节点，在链表中终节点后添加节点
func InsertHeroNode(head *HeroNode, newNode *HeroNode){
	
	temp := head
	//查找链表中终节点
	for {
		if temp.nextNode == nil {
			break
		}
		temp = temp.nextNode
	}
	//添加节点
	temp.nextNode = newNode
}

//有序插入
func InsertHeroNodeOrder(head *HeroNode, newNode *HeroNode){
	
	temp := head
	flag := true
	//
	for {

		if temp.nextNode == nil {
			break
		}else if temp.nextNode.id > newNode.id {//控制大小顺序
			break
		}else if temp.nextNode.id == newNode.id {		
			flag = false
			break
		}
		temp = temp.nextNode
	}
	if !flag{
		fmt.Println("编号已存在")
		return
	}else{
		//把原链中的节点指向新加节点的next
		newNode.nextNode = temp.nextNode
		//添加节点
		temp.nextNode = newNode
	}

	
}

//删除节点
func DeleteHeroNode(head *HeroNode, id int){
		
	temp := head
	flag := false
	//
	for {

		if temp.nextNode == nil {
			break
		}else if temp.nextNode.id == id {		
			flag = true
			break
		}
		temp = temp.nextNode
	}

	if flag {
		temp.nextNode = temp.nextNode.nextNode
	}else{
		fmt.Println("要删除ID不存在")
	}
	
}
// 显示链表信息
func ShowLinkedList(head *HeroNode)  {
	temp := head
	if temp.nextNode == nil {//空链表
		return
	}
	for {
		temp = temp.nextNode
		fmt.Printf("链表节点：%v ==>",temp)
		if temp.nextNode == nil {
			break
		}
	}
	fmt.Println()
}

func SingleLinkedList()  {
	//1、创建头结点
	head := &HeroNode{}
	//2、创建节点
	hero1 := &HeroNode{
		id : 1,
		name : "宋江",
		nickName : "及时雨",
	}
	
	hero2 := &HeroNode{
		id : 2,
		name : "卢俊义",
		nickName : "玉麒麟",
	}
	
	hero3 := &HeroNode{
		id : 3,
		name : "吴用",
		nickName : "智多星",	
	}
	//无序
	// InsertHeroNode(head,hero1)
	// InsertHeroNode(head,hero2)
	// InsertHeroNode(head,hero3)
	//有序
	InsertHeroNodeOrder(head,hero2)
	InsertHeroNodeOrder(head,hero3)
	InsertHeroNodeOrder(head,hero1)
	
	ShowLinkedList(head)
	DeleteHeroNode(head,1)
	ShowLinkedList(head)
}