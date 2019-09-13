package src

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	maps := make(map[int]int, 0)
	result := make([]int, 0)
	for _, v := range arr1 {
		maps[v]++
	}
	for _, v := range arr2 {
		if times, ok := maps[v]; ok {
			for i := 0; i < times; i++ {
				result = append(result, v)
			}
			delete(maps, v)
		}
	}
	keys := make([]int, 0)
	for v := range maps {
		keys = append(keys, v)
	}
	sort.Ints(keys)
	for _, v := range keys {
		times := maps[v]
		for i := 0; i < times; i++ {
			result = append(result, v)
		}
	}

	return result
}
func relativeSortArray2(arr1 []int, arr2 []int) []int {
	// 标记已经被占用的下标
	arr1PlaceHolderIndex := 0
	for i := 0; i < len(arr2); i++ {
		for j := arr1PlaceHolderIndex; j < len(arr1); j++ {
			if arr2[i] == arr1[j] {
				arr1[arr1PlaceHolderIndex], arr1[j] = arr1[j], arr1[arr1PlaceHolderIndex]
				arr1PlaceHolderIndex++
			}
		}
	}
	sortInts := arr1[arr1PlaceHolderIndex:]
	sort.Ints(sortInts)
	return append(arr1[:arr1PlaceHolderIndex], sortInts...)
}

// 对arr1中元素排序，符合arr2中的顺序，没出现在arr2的，就按照升序排
func relativeSortArray3(arr1 []int, arr2 []int) []int {
	pre := make(map[int]int)
	count := 0
	for i := 0; i < len(arr2); i++ {
		if _, ok := pre[arr2[i]]; !ok {
			pre[arr2[i]] = count
			count++
		}
	}
	//n := len(arr1)
	sort.Slice(arr1, func(i, j int) bool {
		v_i, ok_i := pre[arr1[i]]
		v_j, ok_j := pre[arr1[j]]
		if ok_i && ok_j {
			return v_i < v_j
		}
		if ok_i && !ok_j {
			return true
		}
		if !ok_i && ok_j {
			return false
		}
		return arr1[i] < arr1[j]

	})
	return arr1
}
