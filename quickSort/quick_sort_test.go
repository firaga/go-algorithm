package quickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	//in := []int{1, 4, 3, 9, 3}
	//in := []int{3, 1, 4, 100}
	//in := []int{3, 4, 2, 9}
	in := []int{-1, -1, 3, 4, 2, 0, 9}
	//in := []int{9, 5, 4, 3, 1}
	quickSort(in)
	fmt.Println(in)
}
