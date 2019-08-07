package src

import "fmt"

func InsertSort(arr []int) {
	fmt.Println("交换前数组=", arr)
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1

		//从大到小排序
		for insertIndex >= 0 && insertVal > arr[insertIndex] {
			arr[insertIndex+1] = arr[insertIndex] //插入值较小，原值后移
			insertIndex--
		}

		//插入
		if insertIndex+1 != i { //如果插入者小，则不添加
			arr[insertIndex+1] = insertVal
		}
	}
	fmt.Println("交换后数组=", arr)

}

func InsertSort1(arr []int) {
	fmt.Println("交换前数组=", arr)

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] { // 如果待排序列中的元素大于等于有有序序列中的元素，则插入
				temp := arr[j-1]
				arr[j-1] = arr[j]
				arr[j] = temp
			} else { // 不需要交换
				break
			}
		}
		fmt.Printf("第%d次插入排序后array=%v \n", i+1, arr)
	}
	fmt.Println("交换后数组=", arr)

}
