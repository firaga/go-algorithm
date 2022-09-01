package binary_tree

import "testing"

func TestPreOrder(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	root := LevelOrderBuild(input)
	PreOrder(root)
}

func TestPreOrderNonRecursion(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	root := LevelOrderBuild(input)
	PreOrderNonRecursion(root)
}
