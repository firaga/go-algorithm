package quickSort

func headSort(arr []int) {
	buildHeap(arr, len(arr))
	heapSize := len(arr)
	for heapSize > 1 { //heapSize ==1 时不需要继续执行,因为--后堆的大小为0
		swap(arr, 0, heapSize-1)
		heapSize--
		heapify(arr, heapSize, 0)
	}
}

func buildHeap(arr []int, len int) {
	lastNoLeafNodeIndex := len/2 - 1
	for i := lastNoLeafNodeIndex; i >= 0; i-- {
		heapify(arr, len, i)
	}
}

//n length of arr
///i  root of a tree or subtree, it's a index of arr
func heapify(arr []int, n, i int) {
	for true {
		maxPos := i
		//判断左子节点
		if i*2+1 < n && arr[maxPos] < arr[i*2+1] {
			maxPos = i*2 + 1
		}
		//判断右子节点
		if i*2+2 < n && arr[maxPos] < arr[i*2+2] {
			maxPos = i*2 + 2
		}
		if maxPos == i {
			break
		}
		swap(arr, i, maxPos)
		//开启下一轮交换
		i = maxPos
	}
}

func swap(arr []int, a int, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}
