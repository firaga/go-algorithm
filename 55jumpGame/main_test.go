package _5jumpGame

import (
	"fmt"
	"testing"
)

func TestLongestJump(t *testing.T) {
	arr := []int{3, 2, 1, 0, 4}
	rs := canJump2(arr)
	fmt.Println(rs)
}

//最远距离
// 还有一种方法是最早开始,从后开始遍历, 判断last节点是否可以到下标0
func canJump2(nums []int) bool {
	var reach int
	reach = 0
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > reach {
			return false
		}
		if i+nums[i] > reach {
			reach = i + nums[i]
		}
	}
	return true
}

func canJump(nums []int) bool {
	res := dfs(0, nums)
	return res
}

//回溯法
func dfs(index int, nums []int) bool {
	l := len(nums)
	if index == l-1 {
		return true
	}
	if nums[index] == 0 {
		return false
	}
	if index >= l {
		return false
	}
	for i := 1; i <= nums[index]; i++ {
		res := dfs(index+i, nums)
		if res == true {
			return res
		}
	}
	return false
}
