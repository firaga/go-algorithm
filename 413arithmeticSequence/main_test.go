package _13arithmeticSequence

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type itemT struct {
	in  []int
	out int
}

func TestAS(t *testing.T) {
	test := []itemT{
		{in: []int{1, 2, 3, 4}, out: 3},
		{in: []int{1}, out: 0},
		{in: []int{7, 7, 7, 7}, out: 3},
	}
	for _, item := range test {
		rs := numberOfArithmeticSlices(item.in)
		fmt.Println(rs)
		assert.Equal(t, item.out, rs)
	}
}
func numberOfArithmeticSlices(nums []int) int {
	return doublePtr(nums)
}

func violence(nums []int) int {
	l := len(nums)
	var res int
	for i := 0; i < l-2; i++ {
		for j := i + 2; j < l; j++ {
			if isAs(nums, i, j) {
				res++
			}
		}
	}
	return res
}

func isAs(nums []int, s int, e int) bool {

	for i := s; i < e-1; i++ {
		if nums[i+1]*2 != nums[i]+nums[i+2] {
			return false
		}
	}
	return true
}

func doublePtr(nums []int) int {
	l := len(nums)
	var res int
	for i := 0; i < l-2; i++ {
		d := nums[i+1] - nums[i]
		for j := i + 2; j < l; j++ {
			if nums[j]-nums[j-1] == d {
				res++
			} else {
				break
			}
		}
	}
	return res
}
