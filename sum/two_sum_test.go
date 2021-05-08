package sum

import (
	"fmt"
	"testing"
)

func TestTowSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := twoSumOld(nums, target)
	fmt.Println(res)
	res2 := twoSum(nums, target)
	fmt.Println(res2)
}
