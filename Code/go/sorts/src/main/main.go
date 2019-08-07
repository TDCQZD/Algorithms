package main

import (
	"fmt"
	
	
)

func main()  {
	// utils.BubbleSort()

	// var arrys =[6] int {1,8, 10, 89, 1000, 1234}  //查找1的下标
	// utils.BinarySeach(&arrys,0,len(arrys)-1,114)
	// var arry =[6] int {1000,1, 89,10, 8, 1234}
	// utils.SelectSortMaxToMin(&arry)
	// utils.SelectSortMinToMax(&arry)
	// utils.InsertSort(arry)
	// utils.InsertSort1(arry)

	arr := [9]int {-9,78,0,23,-567,70, 123, 90, -23}
	fmt.Println("要排序的数组：",arr);
	//调用快速排序
	 utils.QuickSort(0, len(arr) - 1, &arr)
	fmt.Println("排序后的数组：",arr);
	// utils.QuickSort1(&arr,0,len(arr) - 1)
}