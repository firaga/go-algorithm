package quickSort

import (
	"fmt"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{-1, 3, 4, 2, 0, 9}
	//arr := []int{3, 2, 8, 5, 9}
	headSort(arr)
	fmt.Println(arr)
}
func TestTopK(t *testing.T) {
	arr := []int{-1, 3, 4, 2, 0, 9}
	rs := topK(arr, 3)
	fmt.Println(rs)
}
