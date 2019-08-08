package src

import "fmt"

/*TreeNode represent a  tree node*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//前序遍历
func PreOrder(tree *TreeNode) {

	if tree != nil {
		fmt.Println(tree.Val)
		PreOrder(tree.Left)
		PreOrder(tree.Right)
	}
}

//中序遍历
func InfixOrder(tree *TreeNode) {

	if tree != nil {
		InfixOrder(tree.Left)
		fmt.Println(tree.Val)
		InfixOrder(tree.Right)
	}
}

//后序遍历
func PostOrder(tree *TreeNode) {

	if tree != nil {
		PostOrder(tree.Left)
		PostOrder(tree.Right)
		fmt.Println(tree.Val)
	}
}
