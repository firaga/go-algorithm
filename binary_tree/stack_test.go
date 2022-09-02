package binary_tree

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	s.Push(1)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
