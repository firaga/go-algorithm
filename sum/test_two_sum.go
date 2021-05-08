package sum

//1. hash,减少一次遍历
//2  两个值的index不能相等
func twoSumOld(nums []int, t int) []int {
	res := make([]int, 2)
	hash := make(map[int]int)
	//for i, num := range nums {
	//	hash[num] = i
	//}
	for i := 0; i < len(nums); i++ {
		if v, ok := hash[t-nums[i]]; ok {
			if i != v {
				return []int{v, i}
			}
		} else {
			hash[nums[i]] = i
		}
	}
	return res
}

//官方题解
//题目中有"数组中同一个元素在答案里不能重复出现。" 按下标方式需要判断下标不相等;按值的方式多个下标保存最后一个,但是符合题目要求
//题目中有"你可以按任意顺序返回答案",所以可以循环和赋值放到一个循环里
//值的理解比下标直观一些
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}
