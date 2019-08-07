package utils

import (
	"os"
	"errors"
	"fmt"
)
/*数组实现队列*/
type Queue struct{
	maxSize  int
	array [2]int
	front int
	rear int 
}

//添加数据到队列
func (this *Queue)AddQueue(val int)(err error)  {
	
	//判断队列是否满
	if	this.rear == this.maxSize - 1 {
		err = errors.New("queue full")
		return
	}
	//rear后移，添加元素
	this.rear++
	this.array[this.rear] = val
	return
}
//从队列取出数据
func (this *Queue)AcceptQueue()(value int, err error)  {
	
	
	
	//判断队列是否为空
	if	 this.front  == this.rear  {
		err = errors.New("queue empty")
		return
	}

	fmt.Println("front：",this.front)
	fmt.Println("rear：",this.rear)
	//front后移，取出元素
	this.front++
	value = this.array[this.front]
	
	return
}


//显示队列
func (this *Queue)ShowQueue()  {
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("队列元素array[%d] = %d \n", i, this.array[i])
	}
}

func QueueDemo(){
	
	MainMenu()
}

//显示主菜单
func MainMenu() {
	var key int
	queue := Queue{
		maxSize : 2 ,
		front : -1,
		rear : -1,
	}
	for {
		fmt.Println("-----------------队列操作-----------------")
		fmt.Println("\t\t1 队列添加")
		fmt.Println("\t\t2 队列取值")
		fmt.Println("\t\t3 显示队列")
		fmt.Println("\t\t4 退出系统")
		fmt.Println("------------------------------------------------------")
		fmt.Println("请选择(1-4)：")
		//为了接收用户的输入选择，我们需要定义一个变量
		fmt.Scanln(&key)
		switch key {
			case 1 :
				
				var  id int
				fmt.Print("请输入值:\n")
				fmt.Scanln(&id)
				err := queue.AddQueue(id)
				if err != nil {
					fmt.Println(err)
				}else{
					fmt.Println("ok")
				}
			case 2 :
				val, err := queue.AcceptQueue()
				if err != nil {
					fmt.Println(err)
				}else{
					fmt.Println("ok,取出队列值val=",val)
				}
			case 3 :
				fmt.Println("队列列表：")
				queue.ShowQueue()
		        
			case 4 :
				os.Exit(0)
				fmt.Println("你退出了...")
			default :
				fmt.Println("你的输入有误，请重新选择..")
		}
	}

}
