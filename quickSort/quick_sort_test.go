package quickSort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	input := [][]int{
		//{1, 4, 3, 9, 3},
		{3, 4, 9, 9, 5},
	}
	for _, ints := range input {
		in := make([]int, len(ints))
		copy(in, ints)
		quickSort(in)
		fmt.Println(in)

		//in = make([]int, len(ints))
		//copy(in, ints)
		//quickSort(in)
		//fmt.Println(in)
	}
	//in := []int{1, 4, 3, 9, 3}
	//in := []int{3, 1, 4, 100}
	//in := []int{3, 4, 2, 9}
	//in := []int{-1, -1, 3, 4, 2, 0, 9}
	//in := []int{-1,  3, 4, 2, 0, 9}
	//in := []int{3, 4, 2, 9, 5}
	//quickSort(in)
	//quickSort1_1(in)
	//fmt.Println(in)
	//quickSort2(in2)
	//fmt.Println(in2)
	//in3 := []int{-1, 3, 4, 2, 0, 9}
	//quickSort3(in3, 0, len(in3)-1)
	//fmt.Println(in3)
}
