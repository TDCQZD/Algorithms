package utils

import (
	"fmt"
)

/*冒泡排序
思想：
基本处理思想是通过对相邻两个数据的比较及其交换来达到排序的目的。 
首先，将 n 个元素中的第一个和第二个进行比较，如果两个元素的位置为逆序，则交换两个元素的位置；
进而比较第二个和第三个元素关键字，如此类推，直到比较第 n-1 个元素和第 n 个元素为止； 
上述过程描述了起泡排序的第一趟排序过程，在第一趟排序过程中，我们将关键字最大的元素通过交换操作放到了具有 n 个元素的序列的最一个位置上。 
然后进行第二趟排序，在第二趟排序过程中对元素序列的前 n-1 个元素进行相同操作，其结果是将关键字次大的元素通过交换放到第 n-1 个位置上。一般来说，第 i 趟排序是对元素序列的前 n-i+1 个元素进行排序，使得前 n-i+1 个元素中关键字最大的元素被放置到第 n-i+1 个位置上。排序共进行 n-1 趟，即可使得元素序列按关键字有序。

总结冒泡排序的思路：
1. 每一次的外部排序，确定一个数的顺序
2. 外部排序的次数是 数组的大小-1
3. 每次外部排序的比较次数在减少.
4. 每外部排序的比较次数-1

*/

func BubbleSort()  {
	var arrys =[] int {24,69,80,57,30}
	flag :=false
	for i := 0; i < len(arrys) - 1; i++ {

		for j := 0; j < len(arrys)-1-i; j++ {
			if arrys[j] > arrys[j+ 1] {
				temp := arrys[j]
				arrys[j] = arrys[j+ 1]
				arrys[j + 1] = temp
				flag = true
			}
			fmt.Printf("第%d次外部排序中第%d次内部排序后array=%v \n",i+1,j+1,arrys)
		}
		fmt.Printf("第%d次外部排序后array=%v \n" ,i+1,arrys)
		if flag {
			flag = false
		}else {
			break
		}
	}
	fmt.Println("排序后的数组arry=",arrys)

}

/*二分查找
思想：

思路：
1. 先找到这个 数组的中间的下标  
midIndex =  (leftIndex + rightIndex) / 2
2.  开始比较  让 findVal  和  arr[midIndex ]  比较
3.  如果  findVal  >  arr[midIndex ]   , 则应当到 arr数组 [  midIndex + 1 ~~~~ rightIndex  ]
4.  如果  findVal  <  arr[midIndex ]   , 则应当到 arr数组 [  leftIndex ~~~~ midIndex -1   ]
5.  如果  findVal  ==  arr[midIndex ]   , 找到
6.  将 2, 3, 4, 5 的逻辑进行递归.
7.  什么情况下，表示找不到 当   leftIndex   > rightIndex  成立，则表示没有

*/
func BinarySeach(arr *[6]int,startIndex int,lastIndex int,searchValue int)  {
	
	if startIndex > lastIndex {
		fmt.Println("对不起，找不到")
		return 
	}

	middleIndex :=( startIndex + lastIndex ) / 2

	if searchValue < arr[middleIndex] {
       BinarySeach(arr,startIndex,middleIndex - 1,searchValue)
	}else if searchValue > arr[middleIndex] {
		BinarySeach(arr,middleIndex + 1,lastIndex,searchValue)
	}else{
		fmt.Printf("二分查找的值%d的下标是%d \n",searchValue,middleIndex)
	}
}
/*选择排序
算法原理
接选择排序的第一趟处理是从数据序列所有n个数据中选择一个最小的数据作为有序序列中的第1个元素并将它定位在第一号存储位置，
第二趟处理从数据序列的n-1个数据中选择一个第二小的元素作为有序序列中的第2个元素并将它定位在第二号存储位置，
依此类推，当第n-1趟处理从数据序列的剩下的2个元素中选择一个较小的元素作为有序序列中的最后第2个元素并将它定位在倒数第二号存储位置，
至此，整个的排序处理过程就已完成。

思想：选择排序（select sorting）也是一种简单的排序方法。
它的基本思想是：第一次从R[0]~R[n-1]中选取最小值，与R[0]交换，
第二次从R[1]~R[n-1]中选取最小值，与R[1]交换，
第三次从R[2]~R[n-1]中选取最小值，与R[2]交换，…，
第i次从R[i-1]~R[n-1]中选取最小值，与R[i-1]交换，…, 
第n-1次从R[n-2]~R[n-1]中选取最小值，与R[n-2]交换，
总共通过n-1次，得到一个按排序码从小到大排列的有序序列。
*/
func SelectSortMaxToMin(arr *[6]int)  {
	fmt.Println("选择排序:从大到小")
	fmt.Println("交换前数组=",arr)
	for i := 0; i < len(arr) - 1 ; i++ {
	 	max := arr[i]
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {	
			// 找最大值
			if max < arr[j]	{
				max = arr[j]
				maxIndex = j
			}	
		}
		//交换
		if maxIndex != i {//如果比较值小于假定最大值，不交换，提高效率
			arr[i], arr[maxIndex] = arr[maxIndex],arr[i]
		}
	}

	fmt.Println("交换后数组=",arr)
}
func SelectSortMinToMax(arr *[6]int)  {
	fmt.Println("选择排序:从到到大")
	fmt.Println("交换前数组=",arr)
	for i := 0; i < len(arr) - 1 ; i++ {
	 	min := arr[i]
		minIndex := i
		for j := i + 1; j < len(arr); j++ {	
			// 找最大值
			if min > arr[j]	{
				min = arr[j]
				minIndex = j
			}	
		}
		//交换
		if minIndex != i {
			arr[i], arr[minIndex] = arr[minIndex],arr[i]
		}
	}

	fmt.Println("交换后数组=",arr)
}

/*插入排序
思想：插入排序（Insertion Sorting）的基本思想是：
把n个待排序的元素看成为一个有序表和一个无序表，
开始时有序表中只包含一个元素，无序表中包含有n-1个元素，
排序过程中每次从无序表中取出第一个元素，把它的排序码依次与有序表元素的排序码进行比较，将它插入到有序表中的适当位置，
使之成为新的有序表。
*/

func InsertSort(arr [6]int)  {
	fmt.Println("交换前数组=",arr)
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i-1

		//从大到小排序
		for insertIndex >= 0 && insertVal  > arr[insertIndex] {
			arr[insertIndex + 1] = arr[insertIndex] //插入值较小，原值后移
			insertIndex--
		}

		//插入
		if insertIndex + 1 != i {//如果插入者小，则不添加
			arr[insertIndex + 1] = insertVal
		}
	}
	fmt.Println("交换后数组=",arr)

}

func InsertSort1(arr [6]int)  {
	fmt.Println("交换前数组=",arr)

	for  i:= 0; i < len(arr) - 1; i++ {
		for  j := i + 1; j > 0; j-- {
			if (arr[j] < arr[j - 1]) {// 如果待排序列中的元素大于等于有有序序列中的元素，则插入
				temp := arr[j - 1];
				arr[j - 1] = arr[j];
				arr[j] = temp;
			} else { // 不需要交换
				break;
			}
		}
	}
	fmt.Println("交换后数组=",arr)

}
/*快速排序
思想：
快速排序法介绍:
快速排序（Quicksort）是对冒泡排序的一种改进。
基本思想是：通过一趟排序将要排序的数据分割成独立的两部分，
其中一部分的所有数据都比另外一部分的所有数据都要小，
然后再按此方法对这两部分数据分别进行快速排序，整个排序过程可以递归进行，
以此达到整个数据变成有序序列
*/
//快速排序
//说明
//1. left 表示 数组左边的下标
//2. right 表示数组右边的下标
//3  array 表示要排序的数组
func QuickSort(left int, right int, array *[9]int) {
	l := left
	r := right
	// pivot 是中轴， 支点
	pivot := array[(left + right) / 2]
	temp := 0

	//for 循环的目标是将比 pivot 小的数放到 左边
	//  比 pivot 大的数放到 右边
	for ; l < r; {
		//从  pivot 的左边找到大于等于pivot的值
		for ; array[l] < pivot; {
			l++
		}
		//从  pivot 的右边边找到小于等于pivot的值
		for ; array[r] > pivot; {
			r--
		}
		// 1 >= r 表明本次分解任务完成, break
		if l >= r { 
			break
		}
		//交换
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		//优化
		if array[l]== pivot  {
			r--
		}
		if array[r]== pivot {
			l++			
		}
	}
	// 如果  1== r, 再移动下
	if l == r {
		 l++
		 r--
	}
	// 向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	// 向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}



