package src

import (
	"fmt"
)

/*迷宫问题
1、创建二维数组，模拟迷宫
	1：墙
	0：未走过的通道
	2：走过的通路
	3：走过的死路
2、老鼠找路的函数
	1)i和j测试地图的点
	2)结束的条件 mazeMap[6][5]=2 
	3)可以走的点 mazeMap[i][j]=0 
	4)方向上下左右
3、思考题：找最短路径 
*/

func SeachWay(mazeMap *[8][7]int,i int,j int,) bool {
	if mazeMap[6][5] == 2 {//结束条件
		return true
	}
	if mazeMap[i][j] == 0 { //探索未走过的路
		/*
		//假设此路为通，按上下左右探索
		mazeMap[i][j] = 2
		if SeachWay(mazeMap,i - 1,j){//上
			return true
		}else if SeachWay(mazeMap,i + 1,j){//下
			return true
		}else if SeachWay(mazeMap,i ,j - 1){//左
			return true
		}else if SeachWay(mazeMap,i,j + 1){//右
			return true
		}else{//死路
			mazeMap[i][j] = 3
			return false
		}
	
	//假设此路为通，按下右上左探索
		mazeMap[i][j] = 2
		if SeachWay(mazeMap,i + 1,j){//下
			return true
		}else if SeachWay(mazeMap,i,j + 1){//右
			return true
		}else if SeachWay(mazeMap,i - 1,j){//上
			return true
		}else if SeachWay(mazeMap,i ,j - 1){//左
			return true
		}else{//死路
			mazeMap[i][j] = 3
			return false
		}
*/
		//假设此路为通，按上右下左探索
		mazeMap[i][j] = 2
		if SeachWay(mazeMap,i - 1,j){//下
			return true
		}else if SeachWay(mazeMap,i,j + 1){//右
			return true
		}else if SeachWay(mazeMap,i + 1,j){//上
			return true
		}else if SeachWay(mazeMap,i ,j - 1){//左
			return true
		}else{//死路
			mazeMap[i][j] = 3
			return false
		}

	}
	return false
}

func Maze()  {
	// 1、创建二维数组
	var mazeMap [8][7]int
	for i := 0; i < 7; i++ {
		mazeMap[0][i] = 1
		mazeMap[7][i] = 1
		mazeMap[i][0] = 1
		mazeMap[i][6] = 1
		// mazeMap[4][i] = 1	//设置障碍
	}
	//设置障碍

	fmt.Println("迷宫初始化")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mazeMap[i][j],"\t")
		}
		fmt.Println()
	}
	//探索
	SeachWay(&mazeMap,1,1)

	fmt.Println("探索迷宫结束")
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(mazeMap[i][j],"\t")
		}
		fmt.Println()
	}

}