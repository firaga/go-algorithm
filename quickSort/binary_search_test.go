package quickSort

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	inputs := [][]int{
		{3, 4, 5, 6, 7},
	}
	var idx int
	for _, input := range inputs {
		//idx := binarySearch(input, 5)
		//fmt.Println(idx)
		idx = binarySearch(input, 7)
		fmt.Println(idx)
		//idx = binarySearch(input, 2)
		//fmt.Println(idx)
	}
}

func TestBinarySearchInsertPos(t *testing.T) {
	inputs := [][]int{
		{3, 4, 5, 6, 7},
	}
	var idx int
	for _, input := range inputs {

		//idx = binarySearch(input, 7)
		//fmt.Println(idx)
		idx = binarySearchInsertPos(input, 2)
		fmt.Println(idx)
		idx = binarySearchInsertPos(input, 3)
		fmt.Println(idx)
		idx = binarySearchInsertPos(input, 4)
		fmt.Println(idx)
		idx = binarySearchInsertPos(input, 7)
		fmt.Println(idx)
		idx = binarySearchInsertPos(input, 8)
		fmt.Println(idx)

	}
}
