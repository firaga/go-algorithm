package threeSum

import (
	"fmt"
	"testing"
	"time"
)

func TestThreeSum(t *testing.T) {
	//nums := []int{-1, -1, 0, 1, 2. - 1, -4}
	//nums := []int{-1, 0, 1, 2, -1, -4, 2}
	//-4,-1,-1,-1,0,1,2
	//nums := []int{0, 0, 0, 0}
	// -4,-3,-2,-1,-1,0,0,1,2,3,4
	nums := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}

	start := time.Now()
	res := threeSum(nums)
	cost := time.Since(start)
	fmt.Printf("run time %v\n", cost)
	fmt.Println(len(res))
	fmt.Println(res)

	start = time.Now()
	res = officialThreeSum(nums)
	cost = time.Since(start)
	fmt.Printf("run time %v\n", cost)
	fmt.Println(len(res))
	fmt.Println(res)
}
