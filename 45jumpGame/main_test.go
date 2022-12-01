package _5jumpGame

import (
	"fmt"
	"testing"
)

func TestJump(t *testing.T) {
	//arr := []int{2, 3, 0, 1, 4}
	arr := []int{1, 2, 1, 1, 1}
	rs := jump(arr)
	fmt.Println(rs)
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
