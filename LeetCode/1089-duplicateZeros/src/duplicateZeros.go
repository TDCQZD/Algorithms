package src

func duplicateZeros(arr []int)  {
    var ans []int
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			ans = append(ans, 0)
			ans = append(ans, 0)
		} else {
			ans = append(ans, arr[i])
		}
	}
	for i:=0; i< len(arr);i++{
		arr[i]=ans[i]
	}
}