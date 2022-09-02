package binary_tree

import "testing"

func TestInOrder(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	root := LevelOrderBuild(input)
	InOrder(root)
}
func TestInOrderNonRecursive(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	root := LevelOrderBuild(input)
	InOrderNonRecursive(root)
}
