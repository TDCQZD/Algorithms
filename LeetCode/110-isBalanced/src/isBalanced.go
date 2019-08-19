package src

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if nil == root {
		return true
	}
	diff := abs(depth(root.Left) - depth(root.Right))
	if diff > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if nil == root {
		return 0
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
