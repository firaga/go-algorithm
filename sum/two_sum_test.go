package sum

import (
	"fmt"
	"testing"
)

func TestTowSum(t *testing.T) {
	nums := []int{3, 3, 3, 4}
	target := 7
	res := twoSumOld(nums, target)
	fmt.Println(res)
	res2 := twoSumOld(nums, target)
	fmt.Println(res2)
}
