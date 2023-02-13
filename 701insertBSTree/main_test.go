package _01insertBSTree

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TestInsert(t *testing.T) {
	tree := LevelOrderBuild([]int{4, 2, 7, 1, 3})
	newTree := insertIntoBST(tree, 5)
	fmt.Println(newTree)
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = new(TreeNode)
		root.Val = val
		return root
	}
	if val < root.Val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

func LevelOrderBuild(input []int) *TreeNode {
	return levelOrderBuild(input, 0)
}

func levelOrderBuild(input []int, i int) *TreeNode {
	if len(input) == 0 || i >= len(input) || input[i] == 0 {
		return nil
	}
	node := &TreeNode{
		Val:   input[i],
		Left:  nil,
		Right: nil,
	}
	node.Left = levelOrderBuild(input, 2*i+1)
	node.Right = levelOrderBuild(input, 2*i+2)

	return node
}
