package _queen

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var result []int
	//数组append时可以不make,下标复制注意不能越界,append方式不用管越界,自动扩容
	result = make([]int, 101)
	result[2] = 1
	result[100] = 2
	for i, i2 := range result {
		fmt.Println(i, i2)
	}
}

func TestNQueen(t *testing.T) {
	ns := []int{1, 4}
	for _, n := range ns {
		res := solveNQueens(n)
		fmt.Println(res)
		fmt.Println(len(res))
	}
}

var solutions [][]string

func solveNQueens(n int) [][]string {
	queenPos := make([]int, n)
	for i, _ := range queenPos {
		queenPos[i] = -1
	}
	nQueen(n, 0, queenPos)
	return solutions
}

//递归实现
func nQueen(n int, row int, queenPos []int) {
	if n == row {
		solutions = append(solutions, getResult(queenPos))
		return
	}
	for column := 0; column < n; column++ {
		if isValid(n, row, column, queenPos) {
			queenPos[row] = column
			nQueen(n, row+1, queenPos)
			queenPos[row] = -1
		}
	}
}
func isValid(n, row int, column int, queenPos []int) bool {
	leftUp := column - 1
	rightUp := column + 1
	for i := row - 1; i >= 0; i-- {
		if queenPos[i] == column {
			return false
		}
		if leftUp >= 0 {
			if queenPos[i] == leftUp {
				return false
			}
		}
		if rightUp < n {
			if queenPos[i] == rightUp {
				return false
			}
		}
		leftUp--
		rightUp++
	}
	return true
}

func getResult(queenPos []int) []string {
	n := len(queenPos)
	res := make([]string, n)
	for i := 0; i < n; i++ {
		//rowStr := make([]byte, n)
		var rowStr []byte
		for j := 0; j < n; j++ {
			if queenPos[i] == j {
				rowStr = append(rowStr, 'Q')
			} else {
				rowStr = append(rowStr, '.')
			}
		}
		//res = append(res, string(rowStr))
		res[i] = string(rowStr)
	}
	return res
}
