package src

import (
	"math"
	"sort"
	"strings"
)

func commonChars(A []string) []string {
	length := len(A)
	dist := make(map[int32][]int)
	for index, str := range A {
		for _, s := range str {
			if _, ok := dist[s]; ok {
				dist[s][index]++
			} else {
				dist[s] = make([]int, length)
				dist[s][index] = 1
			}
		}
	}
	result := []string{}
	for key, value := range dist {
		sort.Ints(value[:])
		for index := 0; index < value[0]; index++ {
			result = append(result, string(key))
		}
	}
	return result
}
func commonChars2(A []string) []string {

	res := []string{}
	seen := map[rune]bool{}
	for _, k := range A[0] {
		if seen[k] {
			continue
		}
		seen[k] = true
		n := math.MaxInt32
		for _, w := range A {
			n = min(n, strings.Count(w, string(k)))
		}
		for i := 0; i < n; i++ {
			res = append(res, string(k))
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
func commonChars3(A []string) []string {
	maps:=map[int32]int{}		//maps声明，主要放要返回的字符，及个数
	for _,v:=range A[0]{		//第一个值放到默认maps里面
		maps[v]++
	}
	for i:=1;i< len(A);i++{
        temp:=map[int32]int{}		//每个字符串里面的字符都放到temp里面，用于与maps的比较
		for _,data:=range A[i]{
			if maps[data]>0{		//如果maps存在的话计数，不存在的话不是有效数据
				temp[data]++
			}
		}
		for k,_:=range maps{
			if temp[k]==0{			//如果temp里面不存在，表示maps里面有无效数据
				delete(maps,k)
			}else if maps[k]>temp[k]{	//如果temp里面的数据更少，更新maps数据
				maps[k]=temp[k]
			}
		}
	}
	A=nil
	for k,v:=range maps {		//封装返回值
		for v>0{
			A=append(A, string(k))
			v--
		}
	}
	return A
}
