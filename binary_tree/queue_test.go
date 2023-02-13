package binary_tree

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()
	queue.Add(1)
	queue.Add(2)
	queue.Add(3)
	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
}
