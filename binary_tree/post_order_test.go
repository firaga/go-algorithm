package binary_tree

import "testing"

func TestPostOrder(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	root := LevelOrderBuild(input)
	PostOrder(root)
}

func TestPostOrderNonRecursive(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	root := LevelOrderBuild(input)
	PostOrderNonRecursive(root)
}
