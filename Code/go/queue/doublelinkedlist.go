package utils

import (
	"fmt"
)

/*链表实现队列*/

/*双链表*/

type HeroNodeD struct{
	//数据本身
	id int
	name string
	nickName string
	//nextNode指向一个和它本身存储指向下一个节点的指针
	nextNode *HeroNodeD
	previousNode *HeroNodeD 
}


func InsertHeroNodeD(head *HeroNodeD, newNode *HeroNodeD){
	
	temp := head
	//查找链表中终节点
	for {
		if temp.nextNode == nil {
			break
		}
		temp = temp.nextNode
	}
	//添加节点
	temp.nextNode = newNode.previousNode
	
}

//有序插入
func InsertHeroNodeOrderD(head *HeroNodeD, newNode *HeroNodeD){
	
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
		newNode.previousNode = temp
		if temp.nextNode != nil {
		temp.nextNode.previousNode =newNode
		temp.nextNode = newNode
		}
		
	}

	
}

//删除节点
func DeleteHeroNodeD(head *HeroNodeD, id int){
		
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
		if temp.nextNode != nil{
		temp.nextNode.previousNode = temp 
		}
		
	}else{
		fmt.Println("要删除ID不存在")
	}
	
}
// 显示链表信息 ——正序
func ShowLinkedListD(head *HeroNodeD)  {
	temp := head
	if temp.nextNode == nil {//空链表
		return
	}
	for {
		temp = temp.previousNode
		fmt.Printf("链表节点：%v ==>",temp)
		if temp.previousNode == nil {
			break
		}
	}
	fmt.Println()
}

//显示链表信息 ——逆序
func ShowLinkedListDR(head *HeroNodeD)  {
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






func DoubleLinkedList()  {
	//1、创建头结点
	head := &HeroNodeD{}
	// body := &HeroNodeD{}
	//2、创建节点
	hero1 := &HeroNodeD{
		id : 1,
		name : "宋江",
		nickName : "及时雨",
	}

	hero2 := &HeroNodeD{
		id : 2,
		name : "卢俊义",
		nickName : "玉麒麟",
	}

	hero3 := &HeroNodeD{
		id : 3,
		name : "吴用",
		nickName : "智多星",	
	}
	InsertHeroNodeD(head,hero1)
	InsertHeroNodeD(head,hero2)
	InsertHeroNodeD(head,hero3)
	ShowLinkedListD(head)
}