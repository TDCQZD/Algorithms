package utils

import (
	"fmt"
)


type Hero struct{
	ID int
	Name string 
	Left *Hero
	Right *Hero
}
//前序遍历
func PreOrder(hero *Hero)  {
	
	if hero != nil{
		fmt.Printf("ID=%d,name=%s\n",hero.ID,hero.Name)
		PreOrder(hero.Left)
		PreOrder(hero.Right)
	}
}

//中序遍历
func InfixOrder(hero *Hero)  {
	
	if hero != nil{
		InfixOrder(hero.Left)
		fmt.Printf("ID=%d,name=%s\n",hero.ID,hero.Name)
		InfixOrder(hero.Right)
	}
}

//后序遍历
func PostOrder(hero *Hero)  {
	
	if hero != nil{
		PostOrder(hero.Left)
		PostOrder(hero.Right)
		fmt.Printf("ID=%d,name=%s\n",hero.ID,hero.Name)
	}
}

//二叉树遍历
func TreeTraversing()  {
	
	root := &Hero{
		ID : 1,
		Name : "宋江",
	}
	rootLeft := &Hero{
		ID : 2,
		Name : "吴用",
	}
	rootRirgt := &Hero{
		ID : 3,
		Name : "卢俊义",
	}
	root.Left = rootLeft
	root.Right =  rootRirgt

	left := &Hero{
		ID : 5,
		Name : "公孙胜",
	}

	right2 := &Hero{
		ID : 6,
		Name : "柴秀",
	}
	
	rootLeft.Left = left
	rootLeft.Right = right2
	right := &Hero{
		ID : 4,
		Name : "林冲",
	}
	rootRirgt.Right = right

	fmt.Println("前序遍历")
	PreOrder(root)
	fmt.Println("中序遍历")
	InfixOrder(root)
	fmt.Println("后序遍历")
	PostOrder(root)
	
}