package sum

func twoSum(nums []int, t int) []int {
	res := make([]int, 2)
	hash := make(map[int]int)
	for i, num := range nums {
		hash[num] = i
	}
	for i := 0; i < len(nums); i++ {
		if v, ok := hash[t-nums[i]]; ok {
			if i != v {
				return []int{i, v}
			}
		}
	}
	return res
}
