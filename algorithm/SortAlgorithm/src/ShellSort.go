package src

func ShellSort(arr []int) []int {
	L := len(arr)

	for gap := L / 2; gap > 0; gap /= 2 {
		for i := gap; i < L; i++ {
			insert(arr, gap, i)
		}
	}

	return arr
}

func insert(arr []int, gap, i int) {
	tmp := arr[i]
	j := i - gap
	for j >= 0 && tmp < arr[j] {
		arr[j+gap] = arr[j]
		j -= gap
	}
	arr[j+gap] = tmp
}
