package _299maxLeft

import (
	"fmt"
	"testing"
)

func replaceElements(arr []int) []int {
	j := -1
	var temp int
	for i := len(arr) - 1; i >= 0; i-- {
		temp = arr[i]
		arr[i] = j
		if temp > j {
			j = temp
		}
	}
	return arr
}

func TestReplaceElement(t *testing.T) {
	in := []int{17, 18, 5, 4, 6, 1}
	res := replaceElements(in)
	fmt.Println(res)
}
