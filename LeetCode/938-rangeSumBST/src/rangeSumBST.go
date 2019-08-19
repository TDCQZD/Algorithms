package src

//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	ans := make([]int, 1)
	if nil != root {
		dfs(root, L, R, ans)
	}
	return ans[0]
}
func dfs(node *TreeNode, L int, R int, res []int) {

	if nil != node {
		if L <= node.Val && node.Val <= R {
			res[0] += node.Val
		}
		if L < node.Val {
			dfs(node.Left, L, R, res)
		}
		if node.Val < R {
			dfs(node.Right, L, R, res)
		}
	}
}
