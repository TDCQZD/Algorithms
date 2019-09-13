package src

func removeDuplicates(nums []int) int {
    if nil == nums || len(nums) < 1 {
		return 0
	}

	i := 2
	for  i < len(nums) {
		if nums[i] == nums[i-2] {	// 删除第i个元素
			nums=append(nums[:i],nums[i+1:]...)
		}else{
			i++
		}
	}

	return i 
}
func removeDuplicates1(nums []int) int {
	if len(nums) <= 1{
		return  len(nums)
	}
	
	 curr:= 1          //新数组中有效位置的最后一位，新加入的数据应当写到current+1
	for  i:= 2; i < len(nums);i++ {//从第三位开始循环，前两位无论如何都是要加入新数组的{
		if (nums[i] != nums[curr - 1]){  //符合条件，加入新数组
			curr++
			nums[curr] = nums[i]
		}
	}
	return curr+1


}