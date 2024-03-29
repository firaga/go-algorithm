package _2pickWater

import (
	"fmt"
	"testing"
)

func TestWater(t *testing.T) {
	//[0,1,0,2,1,0,1,3,2,1,2,1]
	//height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	height := []int{4, 2, 0, 3, 2, 5}
	res := trapMax(height)
	fmt.Println(res)
}

func trapMax(height []int) int {
	l := len(height)
	leftMaxSlice := make([]int, l)
	rightMaxSlice := make([]int, l)
	var leftMax, rightMax int
	leftMaxSlice[0] = height[0]
	leftMax = height[0]
	for i := 1; i < l; i++ {
		//错误点 左侧最高判断元素错误
		leftMaxSlice[i] = max(leftMax, height[i])
		if height[i] > leftMax {
			leftMax = height[i]
		}
	}
	rightMaxSlice[l-1] = height[l-1]
	rightMax = height[l-1]
	for i := l - 2; i >= 0; i-- {
		rightMaxSlice[i] = max(rightMax, height[i])
		if height[i] > rightMax {
			rightMax = height[i]
		}
	}
	var res int
	for i := 0; i < l; i++ {
		res += min(leftMaxSlice[i], rightMaxSlice[i]) - height[i]
	}
	return res
}

func trap(height []int) int {
	l := len(height)
	leftMaxSlice := make([]int, l)
	rightMaxSlice := make([]int, l)
	leftMaxSlice[0] = height[0]
	for i := 1; i < l; i++ {
		//错误点1 左侧最高判断元素错误

		//错误点2 改了还是不对, 没有用上次的动态规划值,而是用了截止目前的全局最大值
		//补充,可以用最大max值,注意需要初始化,另可以用上次值,所以没必要加上max值变量
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
