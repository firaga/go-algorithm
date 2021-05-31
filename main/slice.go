package main

import "fmt"

//https://www.runoob.com/go/go-slice.html
func main() {
	var numbers []int
	numbers = make([]int, 0)

	printSlice(numbers)
	numbers = append(numbers, 3)
	printSlice(numbers)
	numbers = append(numbers, 3)
	printSlice(numbers)

	if numbers == nil {
		fmt.Printf("切片是空的")
	}
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
