package quickSort

func binarySearch(input []int, s int) (idx int) {
	lenth := len(input)

	left := 0
	right := lenth - 1

	for left <= right {
		mid := (right-left)/2 + left
		if s < input[mid] {
			right = mid - 1
		} else if s > input[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func binarySearchInsertPos(nums []int, target int) (idx int) {
	length := len(nums)

	left := 0
	right := length - 1

	for left <= right {
		mid := (right-left)/2 + left
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}
	return left
}
