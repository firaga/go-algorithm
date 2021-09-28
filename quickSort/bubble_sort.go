package quickSort

import "fmt"

func bubbleSort(nums []int) {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				//交换
				nums[j], nums[j-1] = nums[j-1], nums[j]
				count++
			}
		}
	}
	fmt.Println(count)
}

// 网上找的错误代码, 因为第二轮没有忽略已经冒泡出去的数值, 会多比较几次
func bubbleSort2(sli []int) []int {
	count := 0
	len := len(sli)
	//该层循环控制 需要冒泡的轮数
	for i := 1; i < len; i++ {
		//该层循环用来控制每轮 冒出一个数 需要比较的次数
		for j := 0; j < len-1; j++ {
			if sli[i] < sli[j] {
				sli[i], sli[j] = sli[j], sli[i]
				count++
			}
		}
	}
	fmt.Println(count)
	return sli
}
