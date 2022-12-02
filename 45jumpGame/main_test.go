package _5jumpGame

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	//arr := []int{2, 3, 0, 1, 4}
	arr := []int{1, 2, 1, 1, 1}
	rs := jump2(arr)
	fmt.Println(rs)
}

//贪婪,不太会, 多看看
func jump2(nums []int) int {
	length := len(nums)
	end := 0
	maxPosition := 0
	steps := 0
	for i := 0; i < length-1; i++ {
		maxPosition = max(maxPosition, i+nums[i])
		if i == end {
			end = maxPosition
			steps++
		}
	}
	return steps
}

//递归实现
func jump(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			//错误点: j协程了dp[j]
			if j+nums[j] >= i {
				//错误点: 第二个参数写错了dp[j]+nums[j]
				if dp[i] == 0 {
					dp[i] = dp[j] + 1
				} else {
					dp[i] = min(dp[i], dp[j]+1)
				}
			}
		}
	}
	return dp[len(nums)-1]
}

func min(a int, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
