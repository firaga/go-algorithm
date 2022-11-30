package _13robGames2

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	//input := []int{2, 3, 2}
	//input := []int{1, 2, 3, 1}
	input := []int{1, 2, 1, 1}
	res := rob(input)
	fmt.Println(res)
}

func rob(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	if l == 2 {
		return max(nums[0], nums[1])
	}
	//错误点
	return max(helper(nums, 0, l-2), helper(nums, 1, l-1))
}

func helper(nums []int, s int, e int) int {
	a := nums[s]
	b := max(nums[s], nums[s+1])
	var c int
	//错误点 i的初始值
	//错误点 判断条件协程了i<len
	for i := s + 2; i <= e; i++ {
		c = max(a+nums[i], b)
		a = b
		b = c
	}
	return b
}

func max(a int, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}
