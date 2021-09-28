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

func quickSort2(sortArray []int, left, right int) {
	if left < right {
		key := sortArray[(left+right)/2]
		i := left
		j := right

		for {
			for sortArray[i] < key {
				i++
			}
			for sortArray[j] > key {
				j--
			}
			if i >= j {
				break
			}
			sortArray[i], sortArray[j] = sortArray[j], sortArray[i]
		}

		quickSort2(sortArray, left, i-1)
		quickSort2(sortArray, j+1, right)
	}
}

func QuickSort2(values []int) {
	quickSort2(values, 0, len(values)-1)
}
