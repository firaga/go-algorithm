package quickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	in := []int{1, 5, 4, 9, 3, 9}
	//in := []int{9, 5, 4, 3, 1}
	quickSort(in)
	fmt.Println(in)
}
