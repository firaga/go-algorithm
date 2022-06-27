package quickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	//in := []int{1, 4, 3, 9, 3}
	//in := []int{3, 1, 4, 100}
	//in := []int{3, 4, 2, 9}
	//in := []int{-1, -1, 3, 4, 2, 0, 9}
	//in2 := []int{-1,  3, 4, 2, 0, 9}
	//in2 := []int{3, 4, 2, 9}
	//quickSort(in)
	//fmt.Println(in)
	//QuickSort2(in2)
	//fmt.Println(in2)
	in3 := []int{-1, 3, 4, 2, 0, 9}
	QuickSort3(in3, 0, len(in3)-1)
	fmt.Println(in3)
}
