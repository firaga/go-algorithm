package quickSort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 2, 1, 5, 4, 6}
	bubbleSort(arr)
	fmt.Println(arr)
	arr2 := []int{3, 2, 1, 5, 4, 6}
	bubbleSort2(arr2)
	fmt.Println(arr2)
}
