package myTest

import (
	"fmt"
	"testing"
)

//pokekara
//二维数组到左下角最短路径

func TestDoSomething(t *testing.T) {
	//二维数组初始化
	input := make([][]int, 4)
	//for i := 0; i < 4; i++ {
	//input[i] = make([]int, 4)
	//}
	input[0] = []int{1, 3, 5, 9}
	input[1] = []int{8, 1, 3, 4}
	input[2] = []int{5, 0, 6, 1}
	input[3] = []int{8, 8, 4, 0}

	fmt.Println(input)
	res := doSomething(input, 4, 4)
	fmt.Println(res)
}

func doSomething(input [][]int, row int, height int) int {
	//二维map初始化
	mid := make(map[int]map[int]int, row)
	for i := 0; i < row; i++ {
		mid[i] = make(map[int]int)
		//各个位置的正确下标表示
		for j := 0; j < height; j++ {
			if i == 0 && j == 0 { // 左上角
				mid[i][j] = input[i][j]
			} else if i-1 >= 0 && j == 0 { //第一列
				mid[i][j] = mid[i-1][j] + input[i][j]
			} else if i == 0 && j-1 >= 0 { // 第一排
				mid[i][j] = mid[i][j-1] + input[i][j]
			} else if i-1 >= 0 && j-1 >= 0 { //其他位置
				mid[i][j] = min(mid[i-1][j], mid[i][j-1]) + input[i][j]
			}
		}
	}
	fmt.Println(mid)
	return mid[row-1][height-1]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
