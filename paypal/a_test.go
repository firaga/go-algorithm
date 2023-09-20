package paypal

import (
	"fmt"
	"testing"
)

type Param map[string]interface{}

type Show struct {
	Param
}

type T struct{}

func (t T) f(n int) T {
	fmt.Print(n)
	return t
}

func main() {

}

func f(n int) {
	defer fmt.Println(n)
	n += 100
	defer fmt.Println(n)
}

func TestB(t1 *testing.T) {
	f(1)
}

func TestA(t *testing.T) {
	var input [][]int
	doSomething(input)

}

func doSomething(input [][]int) {
	//left := 0                     //left top
	//butom := len(input[0]) - 1    //left buttom
	//right := len(input[0][0]) - 1 // right buttom
	//top := len(input[0][0]) - 1   //left top
	//
	//i := 0
	//j := len(input[0][0]) - 1
	//for {
	//	if j >= left {
	//		//output
	//		j--
	//		if j == left {
	//			left++
	//		}
	//	}
	//	if i >= butom {
	//		//output
	//		i++
	//		if i == butom {
	//			butom--
	//		}
	//	}
	//	if i <= right {
	//		//output
	//		j++
	//		if j == right {
	//			right--
	//		}
	//	}
	//	if i >= top {
	//		//output
	//		i--
	//		if i == top {
	//			top++
	//		}
	//	}
	//	if i == butom && i == top && j == left && j == right {
	//		break
	//	}
	//}
}
