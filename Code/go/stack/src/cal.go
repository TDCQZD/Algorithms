package utils

import (
	"strconv"
	"fmt"
)
/*使用栈来实现综合计算器-自定义优先级[priority]*/
//判断运算符
func (this *Stack)IsOper(value int) bool {
	if value == 42 || value == 43 ||value == 45 || value == 47 {
      return true
	}else{
	return false
	}

}
//计算的方法
func (this *Stack)Cal(n int, m int,oper int) (res int) {
	switch oper {
	case 42:
		res = m * n
	case 43:
		res = m +  n 
	case 45:
		res =  m - n
	case 47:
		res = m / n
	default:
		fmt.Println("输入预算法错误")
	}
	return
}

//返回运算符优先级
func (this *Stack)Priority(oper int)(res int)  {
	if oper == 42 ||  oper == 47 {
		 res = 1
	  } else if oper == 43 ||oper == 45 {
		res = 0
	  }
	  return
}



func StackCalDemo()  {
	//初始化数栈
	numStack := &Stack{
		MaxTop : 20,
		Top : -1,
	}
	//初始化数符号栈
	operaStack := &Stack{
		MaxTop : 20,
		Top : -1,
	}
	//计算表达式
	exp := "3+200*6-2"
	index := 0//扫描exp
	n := 0//数值1
	m := 0//数值2
	oper := 0//运算符
	res := 0//计算值
	keepNum := ""//多位数运算
	for {
		ch := exp[index:index+1]//扫描字符串
		temp := int([]byte(ch)[0])//字符对应的ASCII码
		if operaStack.IsOper(temp){ //符号栈操作
			//1、判断空栈
			if operaStack.Top == -1{//空栈
				operaStack.Push(temp)
			}else{ //非空栈
				top := operaStack.Priority(operaStack.arr[operaStack.Top])
				if top >= operaStack.Priority(temp){ //如果顶栈优先级大于等于
					//出栈运算
					n,_ = numStack.Pop()
					m,_ = numStack.Pop()
					oper,_ = operaStack.Pop()
					res = operaStack.Cal(n,m,oper)
					//结果值入栈
					numStack.Push(res)
					//符号入栈
					operaStack.Push(temp)
					fmt.Println("入栈结果值Priority：",res)
				}else{
					operaStack.Push(temp)
				}
			}
			fmt.Println("入栈操作符：",temp)
			
		}else{//数值栈操作
		
			keepNum +=ch
			//多位数处理
			if index == len(exp)-1{
				val,_:=strconv.ParseInt(keepNum,10,64)
				numStack.Push(int(val))
			}else{
				if operaStack.IsOper(int([]byte(exp[index+1:index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = ""
				} 
			}

	/*
			//单位数运算
			val,_:= strconv.ParseInt(ch,10,64)
			
			numStack.Push(int(val))
			fmt.Println("入栈数据：",int(val))
			*/
		}

		//判断index是否已经扫描到表达式结尾
		if  index +1 ==len(exp) {
			break
		}
		index++

	}

	//扫描入栈结束，数值出栈运算
	for {
		if operaStack.Top == -1{
			break
		}
		n,_ = numStack.Pop()
		m,_ = numStack.Pop()
		oper,_ = operaStack.Pop()
		res = operaStack.Cal(n,m,oper)
		//结果值入栈
		numStack.Push(res)
		fmt.Println("入栈结果值：",res)
	}
	res,_ = numStack.Pop()
	fmt.Printf("%s = %v",exp,res)
}