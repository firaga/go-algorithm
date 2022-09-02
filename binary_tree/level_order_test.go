package binary_tree

import "testing"

func TestLevelOrder(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7}
	root := LevelOrderBuild(input)
	LevelOrder(root)
}
