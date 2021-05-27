package threeSum

import (
	"fmt"
	"testing"
)

func TestThreeSum(t *testing.T) {
	//nums := []int{-1, -1, 0, 1, 2. - 1, -4}
	//nums := []int{-1, 0, 1, 2, -1, -4, 2}
	//-4,-1,-1,-1,0,1,2
	//nums := []int{0, 0, 0, 0}
	// -4,-3,-2,-1,-1,0,0,1,2,3,4
	nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
	res := threeSum(nums)
	fmt.Println(res)
	res = officialThreeSum(nums)
	fmt.Println(res)
}
