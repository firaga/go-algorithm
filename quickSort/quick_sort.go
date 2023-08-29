package quickSort

func quickSort(inputArray []int) {

	//start of partition
	pivot := len(inputArray) - 1
	//遍历指针, 交换位置指针, 指向小于pivot的元素后一个位置
	var traversalPtr, minPrt int
	for ; traversalPtr <= pivot; traversalPtr++ {
		if inputArray[traversalPtr] <= inputArray[pivot] {
			//swap
			inputArray[minPrt], inputArray[traversalPtr] = inputArray[traversalPtr], inputArray[minPrt]
			//递增,指向小于in[pivot]的下一个元素
			//因为为了方便交换最后一个元素，循环使用了《=，包括了自己本身，所以下面递归时需要-1处理
			minPrt++
		}
	}
	//end of partition
	//直接退出
	if len(inputArray) <= 2 {
		return
	}
	//fmt.Println("minPtr", minPrt)
	//if minPrt-2 >= 0 {
	quickSort(inputArray[:minPrt-1])
	quickSort(inputArray[minPrt:])
}

func quickSort11(inputArray []int) {
	//fmt.Println(inputArray)
	pivot := len(inputArray) - 1
	//遍历指针, 交换位置指针, 指向小于pivot的元素后一个位置
	var traversalPtr, minPrt int
	for ; traversalPtr < pivot; traversalPtr++ {
		if inputArray[traversalPtr] <= inputArray[pivot] {
			//swap
			inputArray[minPrt], inputArray[traversalPtr] = inputArray[traversalPtr], inputArray[minPrt]
			//递增,指向小于in[pivot]的下一个元素
			//因为为了方便交换最后一个元素，循环使用了《=，包括了自己本身，所以下面递归时需要-1处理
			minPrt++
		}
	}
	//minPrt指向下一个待使用小于inputArray[pivot]的值，交换povit到此位置，两边分别大于小于此值
	inputArray[minPrt], inputArray[pivot] = inputArray[pivot], inputArray[minPrt]
	//直接退出
	if len(inputArray) <= 2 {
		return
	}
	//fmt.Println("minPtr", minPrt)
	//if minPrt-2 >= 0 {
	quickSort(inputArray[:minPrt])
	quickSort(inputArray[minPrt+1:])
}

//todo 处理重复有问题
//因为是对值比较，重复值存在的话会一直循环，比如{3, 4, 9, 9, 5}，i，j会一直停留在两个9上，重复循环

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

//todo test
func quickSort3(arr []int, left, right int) {
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
	quickSort3(arr, left, l-1)
	quickSort3(arr, l+1, right)
}
