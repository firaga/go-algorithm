package _2pickWater

import (
	"fmt"
	"testing"
)

func TestWater(t *testing.T) {
	//[0,1,0,2,1,0,1,3,2,1,2,1]
	//height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	height := []int{4, 2, 0, 3, 2, 5}
	res := trap(height)
	fmt.Println(res)
}

func trap(height []int) int {
	l := len(height)
	leftMaxSlice := make([]int, l)
	rightMaxSlice := make([]int, l)
	leftMaxSlice[0] = height[0]
	for i := 1; i < l; i++ {
		//错误点1 左侧最高判断元素错误
		//错误点2 改了还是不对, 没有用上次的动态规划值,而是用了截止目前的全局最大值
		leftMaxSlice[i] = max(leftMaxSlice[i-1], height[i])
	}
	rightMaxSlice[l-1] = height[l-1]
	for i := l - 2; i >= 0; i-- {
		rightMaxSlice[i] = max(rightMaxSlice[i+1], height[i])
	}
	var res int
	for i := 0; i < l; i++ {
		res += min(leftMaxSlice[i], rightMaxSlice[i]) - height[i]
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
