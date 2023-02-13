package binary_tree

import (
	"fmt"
	"testing"
)

func TestLevelOrderBuild(t *testing.T) {
	input := []int{1, 2, 3, 0, 0, 6, 7}
	output := LevelOrderBuild(input)
	fmt.Println(output)
}

func TestLevelOrderBuild2(t *testing.T) {
	input := []int{1, 2, 3, 0, 0, 6, 7}
	output := LevelOrderBuild2(input)
	fmt.Println(output)
}
