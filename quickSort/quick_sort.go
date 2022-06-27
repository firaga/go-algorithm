package quickSort

func quickSort(inputArray []int) {

	//fmt.Println(inputArray)
	pivot := len(inputArray) - 1
	//遍历指针, 交换位置指针, 指向小于pivot的元素后一个位置
	var traversalPtr, minPrt int
	for ; traversalPtr <= pivot; traversalPtr++ {
		if inputArray[traversalPtr] <= inputArray[pivot] {
			//swap
			inputArray[minPrt], inputArray[traversalPtr] = inputArray[traversalPtr], inputArray[minPrt]
			//递增,指向小于in[pivot]的下一个元素
			minPrt++
		}
	}
	//直接退出
	if len(inputArray) <= 2 {
		return
	}
	//fmt.Println("minPtr", minPrt)
	//if minPrt-2 >= 0 {
	quickSort(inputArray[:minPrt-1])
	quickSort(inputArray[minPrt:])
}

//todo 处理重复有问题
func quickSort2(sortArray []int, left, right int) {
	if left < right {
		//选定pivot,位置随意
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

//todo test
func QuickSort3(arr []int, left, right int) {
	if left >= right {
		return
	}
	l := left
	r := right
	base := arr[left]
	for l < r {
		for l < r && arr[r] >= base {
			r--
		}
		if l < r {
			arr[l] = arr[r]
		}

		for l < r && arr[l] <= base {
			l++
		}

		if l < r {
			arr[r] = arr[l]
			r++
		}
	}
	arr[l] = base
	QuickSort3(arr, left, l-1)
	QuickSort3(arr, l+1, right)
}
