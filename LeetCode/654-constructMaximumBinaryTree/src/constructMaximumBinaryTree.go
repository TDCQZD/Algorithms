package src

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) <= 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	m := Max(nums)
	root := &TreeNode{Val: nums[m]}
	root.Left = constructMaximumBinaryTree(nums[:m])
	root.Right = constructMaximumBinaryTree(nums[m+1:])
	return root
}

func Max(nums []int) int {
	n := 0
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
			n = i
		}
	}
	return n
}
