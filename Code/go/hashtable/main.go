package main

import (
	"fmt"
	"os"
)

/*使用hash表实现

google公司的一个上机题:

有一个公司,当有新的员工来报道时,要求将该员工的信息加入(性别,年龄,住址..),当输入该员工的id时,要求查找到该员工的 所有信息.

要求: 不使用数据库,尽量节省内存,速度越快越好=>哈希表(散列)

使用链表来实现哈希表, 该链表不带表头[即: 链表的第一个结点就存放雇员信息]
*/

//员工信息结构体
type Emp struct {
	ID       int
	Name     string
	nextNode *Emp
}

func (this *Emp) ShowEmp() {
	fmt.Printf("id= %d的员工在第%d链，信息为%v \n", this.ID, this.ID%7, this)
}

//员工结构体 不带表头，即第一个节点就存放信息
type EmpLink struct {
	Head *Emp
}

//链表添加员工
func (this *EmpLink) InsertLink(emp *Emp) {
	cur := this.Head   //辅助指针,当前链表中第一个节点
	var pre *Emp = nil //辅助指针,cur上一个节点

	/*空链表操作 :直接添加*/
	if cur == nil {
		this.Head = emp
		return
	}

	/*非空链表操作 :找到位置插入
	1、比较cur 和emp,按由小到大顺序
	2、pre在cur前面
	*/

	//1、确定链表中插入的位置
	for {
		if cur != nil { //非空链表
			if cur.ID > emp.ID {
				break
			}
			pre = cur
			cur = cur.nextNode
		} else {
			break
		}
	}
	//2、添加
	pre.nextNode = emp
	emp.nextNode = cur
	/*
		if cur == nil {//链表结尾添加
			pre.nextNode = emp
			emp.nextNode = cur
		}else{
			pre.nextNode = emp
			emp.nextNode = cur
		}
	*/
}

//遍历链表，查找ID
func (this *EmpLink) SearchLinkByID(id int) *Emp {
	if this.Head == nil { //链表为空
		// fmt.Printf("链表 %d 为空\n",num )
		return nil
	}
	cur := this.Head
	for {
		if cur != nil {
			if cur.ID == id {
				fmt.Printf("id=%d 的用户 %v ", id, cur)
				return cur
			} else {
				cur = cur.nextNode
			}
		} else {
			break
		}
	}
	return nil

}
func (this *EmpLink) SearchLink(id int) {
	if this.Head == nil { //链表为空
		// fmt.Printf("链表 %d 为空\n",num )
		return
	}

	cur := this.Head

	for {
		if cur != nil {
			if cur.ID == id {
				fmt.Printf("id=%d 的用户 %v ", id, cur)
				break
			}
			// fmt.Printf("%d = %v ->",num,cur)
			cur = cur.nextNode
		} else {
			break
		}
	}
	fmt.Println()
}

//遍历链表
func (this *EmpLink) ShowLink(num int) {
	if this.Head == nil { //链表为空
		fmt.Printf("链表 %d 为空\n", num)
		return
	}

	cur := this.Head
	fmt.Printf("链表%d =", num)
	for {
		if cur != nil {
			fmt.Printf(" %v ->", cur)
			cur = cur.nextNode
		} else {
			break
		}
	}
	fmt.Println()
}

//HashTable结构体
type HashTable struct {
	LinkArr [7]EmpLink //链表数组
}

//添加员工信息
func (this *HashTable) InsertHash(emp *Emp) {
	// 1、使用散列函数，确定雇员添加的链表
	linkNum := this.HashFunc(emp.ID) //链表下标
	//2、添加道对应的链表
	this.LinkArr[linkNum].InsertLink(emp)

}

//散列方法获取链表下标
func (this *HashTable) HashFunc(id int) int {
	return id % 7
}

//查找用户
func (this *HashTable) SearchHashEmp(id int) *Emp {
	linkNum := this.HashFunc(id)
	cur := this.LinkArr[linkNum].SearchLinkByID(id)
	return cur
	// for i := 0; i < len(this.LinkArr); i++ {
	// 	this.LinkArr[i].SearchLink(i,id)
	// }

}

//显示HashTable所有雇员
func (this *HashTable) ShowHashAllEmp() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].ShowLink(i)
	}
}

func HashTableDemo() {
	key := 0
	var id int
	var userName string

	var hash HashTable

	for {
		fmt.Println("-----------------雇员管理系统-----------------")
		fmt.Println("\t\t\t1 添加用户")
		fmt.Println("\t\t\t2 显示用户")
		fmt.Println("\t\t\t3 查找用户")
		fmt.Println("\t\t\t4 退出系统")
		fmt.Println("------------------------------------------------------")
		fmt.Println("请选择(1-4)：")
		//为了接收用户的输入选择，我们需要定义一个变量
		fmt.Scanln(&key)
		switch key {
		case 1:
			fmt.Print("请输入用户ID:\n")
			fmt.Scanln(&id)
			fmt.Print("请输入用户姓名:\n")
			fmt.Scanln(&userName)
			emp := &Emp{
				ID:   id,
				Name: userName,
			}
			hash.InsertHash(emp)
		case 2:
			hash.ShowHashAllEmp()
		case 3:
			fmt.Print("请输入用户ID:\n")
			fmt.Scanln(&id)
			emp := hash.SearchHashEmp(id)
			if emp != nil {
				emp.ShowEmp()
			} else {
				fmt.Println("该雇员信息不存在")
			}
		case 4:
			os.Exit(0)
			fmt.Println("你退出了该软件...")
		default:
			fmt.Println("你的输入有误，请重新选择..")
		}

	}

}

func main() {
	HashTableDemo()
}
