package threeSum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	l := len(nums)
	var first, second, third int

	//快排
	sort.Ints(nums)

	for first = 0; first < l; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		reversFirstValue := -nums[first]
		for second = first + 1; second < l; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for third = l - 1; third > second; {
				if third < l-1 && nums[third] == nums[third+1] {
					third--
					continue
				}
				if nums[third]+nums[second] > reversFirstValue {
					//second+third > -first
					third--
				} else if nums[third]+nums[second] == reversFirstValue {
					//second+third = -first
					ans = append(ans, []int{nums[first], nums[second], nums[third]})
					third--
				} else {
					//second+third < -first
					second++
				}
			}
		}
	}
	return ans
}
