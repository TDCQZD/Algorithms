package src

import "fmt"

func BulleSort(arrys []int) {

	for i := 0; i < len(arrys)-1; i++ {

		for j := 0; j < len(arrys)-1-i; j++ {
			if arrys[j] > arrys[j+1] {
				// temp := arrys[j]
				// arrys[j] = arrys[j+1]
				// arrys[j+1] = temp
				arrys[j], arrys[j+1] = arrys[j+1], arrys[j]

			}
			fmt.Printf("第%d次外部排序中第%d次内部排序后array=%v \n", i+1, j+1, arrys)
		}
		fmt.Printf("第%d次外部排序后array=%v \n", i+1, arrys)

	}
	fmt.Println("排序后的数组arry=", arrys)
}

// 优化外循环
func BulleSort_optimization(arrys []int) {

	flag := false
	for i := 0; i < len(arrys)-1; i++ {

		for j := 0; j < len(arrys)-1-i; j++ {
			if arrys[j] > arrys[j+1] {
				// temp := arrys[j]
				// arrys[j] = arrys[j+1]
				// arrys[j+1] = temp
				arrys[j], arrys[j+1] = arrys[j+1], arrys[j]
				flag = true
			}
			fmt.Printf("第%d次外部排序中第%d次内部排序后array=%v \n", i+1, j+1, arrys)
		}
		fmt.Printf("第%d次外部排序后array=%v \n", i+1, arrys)
		// 优化外循环
		if flag { // 如果内循环交换过,
			flag = false
		} else {
			break
		}
	}
	fmt.Println("排序后的数组arry=", arrys)
}
