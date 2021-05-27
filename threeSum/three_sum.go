package threeSum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	l := len(nums)
	var first, second, third int
	third = l - 1
	//快排
	sort.Ints(nums)

	for first = 0; first < l; first++ {
		reversFirstValue := -nums[first]
		for second = first + 1; second < l && nums[second] != nums[second-1]; second++ {
			for third > second {
				if nums[third]+nums[second] > reversFirstValue {
					third--
				} else if nums[third]+nums[second] == reversFirstValue {
					ans = append(ans, []int{nums[first], nums[second], nums[third]})
					break
				} else {
					second++
				}
			}
		}
	}
	return ans
}
