package utils

import (
	"fmt"
)

/*
约瑟夫问题(丢手帕问题)
Josephu  问题
Josephu  问题为：设编号为1，2，… n的n个人围坐一圈，约定编号为k（1<=k<=n）的人从1开始报数，数到m 的那个人出列，它的下一位又从1开始报数，数到m的那个人又出列，依次类推，直到所有人出列为止，由此产生一个出队编号的序列。
提示：用一个不带头结点的循环链表来处理Josephu 
问题：先构成一个有n个结点的单循环链表，然后由k结点起从1开始计数，计到m时，对应结点从链表中删除，然后再从被删除结点的下一个结点又从1开始计数，直到最后一个结点从链表中删除算法结束。

分析：
1、构建链表元素结构体
2、编写函数实现结构体构成单向的循环链表
3、遍历链表
4、实现约瑟夫问题
  1）编写函数实现
  2）使用算法
*/

type Boy struct{
	Num int //编号
	nextBoy *Boy // 
} 

func AddBoy(num int) *Boy {
	
	first := &Boy{}//空节点
	curBoy := &Boy{}
	if num < 1 {
		fmt.Println("num 的值错误")
		return first
	}
	//构建循环链表
	for i := 1; i <= num; i++ {		
		boy := &Boy{
			Num : i,
		}
		if i == 1 {
			first =boy
			curBoy = boy
			curBoy.nextBoy =first
		}else{
			curBoy.nextBoy = boy
			curBoy = boy
			curBoy.nextBoy = first
		}

	}
	return first
}

func ShowLinks(first *Boy) (count int) {

	if first.nextBoy == nil {
		fmt.Println("链表为空")
		return
	}

	curBoy := first
	for  {
		fmt.Println(curBoy)
		count++
		if curBoy.nextBoy == first {
			break
		}
		curBoy = curBoy.nextBoy
		
	}
	  return 
}

func PlayGame(first *Boy,startNum int,countNum int)  {
	//1、判断空链表
	if first.nextBoy == nil {
		fmt.Println("链表为空")
		return
	}
	//2、判断startNum 0<startNum<n
	if startNum < 1 || startNum >ShowLinks(first){
		fmt.Println("startNum错误")
		return
	}
	//3、辅助指针，便于删除
	tmp := first
	//4、确定最后一个指针
	for{
		if tmp.nextBoy == first {
			break
		}
		tmp = tmp.nextBoy		
	}
	//5、让first 移动到 startNo
	for i := 0; i < startNum - 1; i++ {
		first = first.nextBoy
		tmp = tmp.nextBoy
	}
	//6、开始数 countNum, 然后就删除first 指向的小孩
	for{
		for i := 0; i < countNum - 1; i++ {
			first = first.nextBoy
			tmp = tmp.nextBoy	
		}
		//删除
		fmt.Printf("%v淘汰 \n",first)
		first = first.nextBoy
		tmp.nextBoy = first

		if tmp == first {
			break
		}
	}
	fmt.Printf("%v 获胜",first)
}

func ShowBoy()  {
	first := AddBoy(5)
	count := ShowLinks(first)
	fmt.Println(count)
	PlayGame(first,2,3)
}