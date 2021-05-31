package fourSum

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	fmt.Println(nums)
	var first, second, third, fourth int
	length := len(nums)
	for first = 0; first < length; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second = first + 1; second < length; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			//fourth在此循环外赋值,只会减少,不会增加,才是双指针法
			fourth = length - 1
			for third = second + 1; third < length; third++ {
				if third > second+1 && nums[third] == nums[third-1] {
					continue
				}
				for fourth > third && nums[fourth]+nums[third] > target-(nums[first]+nums[second]) {
					//fmt.Printf(" %v, %v, %v, %v\n", nums[first], nums[second], nums[third], nums[fourth])
					fourth--
				}
				if fourth == third {
					break
				}

				if nums[third]+nums[fourth]+nums[first]+nums[second] == target {
					ans = append(ans, []int{nums[first], nums[second], nums[third], nums[fourth]})
				}
			}
		}
	}
	return ans
}
