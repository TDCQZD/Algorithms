package src

import "sort"

func lastStoneWeight(stones []int) int {
	weight := 0
	for i := 0; i < len(stones)-1; i++ {
		sort.Ints(stones)
		weight = stones[len(stones)-1] - stones[len(stones)-2]
		stones[len(stones)-2] = 0
		stones[len(stones)-1] = weight
	}
	return stones[len(stones)-1]

}
