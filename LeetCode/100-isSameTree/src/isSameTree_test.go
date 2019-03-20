package src

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	rootA := &TreeNode{Val: 1}
	nodeA1 := &TreeNode{Val: 2}
	// nodeA2 := &TreeNode{Val: 3}
	rootA.Left = nodeA1
	// rootA.Right = nodeA2

	rootB := &TreeNode{Val: 1}
	// nodeB1 := &TreeNode{Val: 2}
	nodeB2 := &TreeNode{Val: 3}
	// rootB.Left = nodeB1
	rootB.Right = nodeB2

	res := isSameTree(rootA, rootB)
	fmt.Println(res)
}
