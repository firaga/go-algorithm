package sum

import (
	"fmt"
	"testing"
)

func TestTowSum(t *testing.T) {
	nums := []int{3,2,4}
	target := 6
	res := twoSum(nums, target)
	fmt.Println(res)
}
