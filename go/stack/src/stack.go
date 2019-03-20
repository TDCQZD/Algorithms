package utils

import (
	"errors"
	"fmt"
)

/*数组模拟栈使用*/
type Stack struct{
	MaxTop int //栈内存
	Top int //栈顶
	//栈底固定，直接使用Top
	arr [20]int
}
/*入栈*/
func (this *Stack)Push(val int) (err error) {
	//1、判断栈是否满
	if this.Top == this.MaxTop - 1 {
		fmt.Println("Push：stack full")
		return errors.New("Push：stack full")
	}
	//2、入栈
	this.Top++
	this.arr[this.Top] =val
	return
}
/*出栈*/
func (this *Stack)Pop() (val int,err error) {
	//1、判断栈是否为空
	if this.Top == -1{
		fmt.Println("Push：stack empty")
		err = errors.New("Pop：stack empty")
		return
	}
	//2、出栈
	val = this.arr[this.Top]
	fmt.Printf("出栈数据：arr[%d]=%d\n",this.Top,val)
	this.Top--
	return
}
/*栈遍历*/
func (this *Stack)ListStack()  {
	//1、判断栈是否为空
	if this.Top == -1{
		fmt.Println("ListStack：stack empty")
		return
	}
	//2、遍历栈
	for i := this.Top; i>= 0 ; i-- {
		fmt.Printf("栈遍历：arr[%d]=%d\n",i,this.arr[i])
	}
}


func StackDemo()  {
	stack := &Stack{
		MaxTop : 5,
		Top : -1,//当栈顶为-1时，表示空栈
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	stack.ListStack()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	// stack.Pop()
	// stack.Pop()
	// stack.Pop()
	stack.ListStack()
	
}


