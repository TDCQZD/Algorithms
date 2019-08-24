package src

func largestSumAfterKNegations(A []int, K int) int {
	AMap := make(map[int]int, 201) //-100 <= A[i] <= 100,这个范围的大小是201
	for t := 0; t < 201; t++ {     //注意：cap =201,如果没有存放足够的数据len=0
		AMap[t] = 0
	}
	for _, v := range A {
		AMap[v+100]++ //将[-100,100]映射到[0,200]上
	}
	i := 0
	for K > 0 {
		for AMap[i] == 0 { //找到A[]中最小的数字
			i++
		}
		AMap[i]--     //此数字个数-1
		AMap[200-i]++ //其相反数个数+1
		if i > 100 {
			i = 200 - i //若原最小数索引>100,则新的最小数索引应为200-i.(索引即number[]数组的下标)
		}

		K--
	}
	sum := 0
	for j := i; j < len(AMap); j++ { //遍历number[]求和

		sum += (j - 100) * AMap[j] //j-100是数字大小,number[j]是该数字出现次数.
	}

	return sum
}
