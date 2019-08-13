package src

import "fmt"

func SelectSortMaxToMin(arr []int) {
	fmt.Println("选择排序:从大到小")
	fmt.Println("交换前数组=", arr)
	for i := 0; i < len(arr)-1; i++ {
		max := arr[i]
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			// 找最大值
			if max < arr[j] {
				max = arr[j]
				maxIndex = j
			}
		}
		//交换
		if maxIndex != i { //如果比较值小于假定最大值，不交换，提高效率
			arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
		}
	}

	fmt.Println("交换后数组=", arr)
}
func SelectSortMinToMax(arr []int) {
	fmt.Println("选择排序:从小到大")
	fmt.Println("交换前数组=", arr)
	for i := 0; i < len(arr)-1; i++ {
		min := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			// 找最大值
			if min > arr[j] {
				min = arr[j]
				minIndex = j
			}
		}
		//交换
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex], arr[i]
		}
	}

	fmt.Println("交换后数组=", arr)
}
