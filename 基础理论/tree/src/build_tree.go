package src

// 题目：重构树
// 已知一颗二叉树的先序遍历序列为ABCDEFG，中序遍历为CDBAEGF，能否唯一确定一颗二叉树？如果可以，请画出这颗二叉树

/*BuildTree build a tree and return*/
func BuildTree(pre, in []int) *TreeNode {

	if len(in) == 0 {
		return nil
	}

	res := &TreeNode{
		Val: pre[0],
	}

	if len(in) == 1 {
		return res
	}

	idx := indexOf(res.Val, in)

	res.Left = BuildTree(pre[1:idx+1], in[:idx])
	res.Right = BuildTree(pre[idx+1:], in[idx+1:])

	return res
}

//indexOf return index of the tree node
func indexOf(val int, nums []int) int {
	for i, v := range nums {
		if v == val {
			return i
		}
	}

	return 0
}
