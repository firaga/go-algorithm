package quickSort

import "fmt"

func quickSort(in []int) {

	//fmt.Println(in)
	pivot := len(in) - 1
	var traversalPtr, minPrt int
	for ; traversalPtr <= pivot; traversalPtr++ {
		if in[traversalPtr] <= in[pivot] {
			//swap
			in[minPrt], in[traversalPtr] = in[traversalPtr], in[minPrt]
			//递增,指向小于in[pivot]的下一个元素
			minPrt++
		}
	}
	if len(in) <= 2 {
		return
	}
	fmt.Println("minPtr", minPrt)
	//if minPrt-2 >= 0 {
	quickSort(in[:minPrt-1])
	quickSort(in[minPrt:])
}
