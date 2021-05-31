package fourSum

import (
	"fmt"
	"testing"
)

func TestFourSum(t *testing.T) {
	//-3, -2, -1, 0, 0, 1, 2, 3, 4
	//nums := []int{-1, 3, 2, 4, 1, -3, 0, -2, 0}
	//nums := []int{2, 2, 2, 2, 2}
	//[-3,-2,-1,0,0,1,2,3]
	nums := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0
	res := fourSum(nums, target)
	fmt.Println(res)
	res = fourSumOfficial(nums, target)
	fmt.Println(res)
}
