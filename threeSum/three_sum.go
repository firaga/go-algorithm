package threeSum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	l := len(ans)
	var first, second, third int
	third = l - 1
	//快排
	sort.Ints(nums)

	for first = 0; first < l; first++ {
		reversFirst := -first
		for second = first + 1; second < l; second++ {
			for ; third > second; third-- {

			}
		}
	}
	return ans
}
