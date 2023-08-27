package quickSort

import (
	"fmt"
	"testing"
)

// 参考1 数据结构与算法之美 - 28堆和堆排序
//	自上而下建堆，对非叶子节点heapify
// 参考2 https://blog.csdn.net/qq_37874758/article/details/120758589
// 自下而上建堆，
func TestHeapSort(t *testing.T) {
	arr := []int{-1, 3, 4, 2, 0, 9}
	//arr := []int{3, 2, 8, 5, 9}
	headSort(arr)
	fmt.Println(arr)
}
func TestTopK(t *testing.T) {
	arr := []int{-1, 3, 4, 2, 0, 9}
	rs := topK(arr, 3)
	fmt.Println(rs)
}
