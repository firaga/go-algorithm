package quickSort

//https://zhuanlan.zhihu.com/p/354485697

import (
	"fmt"
	"math/rand"
	"time"
)

/**
希尔排序：把切片分成n个batch，对每个batch进行插入排序；然后减小batch，再对每个batch进行插入排序；直到bathc等于1
*/
func shellSort(st []int, batchSize int) {
	if batchSize < 1 {
		return
	}
	// i : 每个batch中的元素所在batch的index， 介于[0, batchSize)
	for i := 0; i < batchSize; i++ {
		// 用到了插入排序
		for j := 1; batchSize*j+i < len(st); j++ { // j: 用来获取每个batch所在的第i个元素，拥有多少个batch
			for n := j; n > 0; n-- {
				pre := batchSize*(n-1) + i
				next := batchSize*n + i
				if st[next] < st[pre] {
					st[next], st[pre] = st[pre], st[next]
				}
			}

		}
	}
	shellSort(st, batchSize/2)
}

func main() {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	list := make([]int, 0)
	for i := 0; i < 100; i++ {
		list = append(list, r.Intn(100))
	}
	shellSort(list, len(list)/2)
	fmt.Println(list)
}
