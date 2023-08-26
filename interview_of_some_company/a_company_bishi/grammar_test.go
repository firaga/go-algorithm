package a_company_bishi

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGrammarOfByteAndUtf8(t *testing.T) {
	var a byte = 0x11 //17
	var b uint8 = a
	var c uint8 = a + b
	test(c)
}

func test(c byte) {
	fmt.Println(c)
}

type T struct {
	n int
}

func TestStructInMap(t *testing.T) {
	//m := make(map[int]T)
	//b := m[0].n
	//数组中的struct，map不能访问
	//m[0].n = 1

}

// 未匹配的case项跳过
func TestCase(t *testing.T) {
	isMatch := func(i int) bool {
		switch i {
		case 1:
			//fallthrough
		case 2:
			return true
		}
		return false
	}
	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))
}

type U struct {
	ls []int
}

func foo(t U) {
	t.ls[0] = 100
}

// 未触发slice扩容情况下，地址不变
func TestAltSlice(t *testing.T) {
	var u = U{ls: []int{1, 2, 3}}
	foo(u)
	fmt.Println(u.ls[0])
}

// 未初始化chan读写，阻塞,抛出panic
// 读关闭的channel，继续读完buffer中元素，读完返回0值
// 写关闭channel，panic
// 关闭未初始化的channel， panic
func TestSelectChan(t *testing.T) {
	var ch chan int
	//v, ok := <-ch
	//fmt.Println(v, ok)
	//return
	select {
	case v, ok := <-ch:
		println(v, ok)
	default:
		println("default")
	}
}

func TestInterface(t *testing.T) {
	//x := interface{}(1)//{interface{}|int} 1
	x := interface{}(nil) //{interface{}}nil
	y := (*int)(nil)      //单纯类型的nil，不涉及interface结构
	a := y == x           //x,y的类型不同
	b := y == nil
	d, c := x.(interface{}) //不知道原因，可能是eface断言必须_type有值;或者eface _type=nil不能断言为interface{}
	fmt.Printf("%v,%T", x, x)
	println(a, b, c, d)

	h := interface{}(3) //{interface|int} 3
	//i, j := h.(interface{}) //i:={interface|int} 3
	i, j := h.(int) //i:={int} 3
	fmt.Println(h, i, j)
}

func TestMultiAssign(t *testing.T) {
	i := 1
	s := []string{"A", "B", "C"}
	i, s[i-1] = 2, "Z" //编译器自动生成临时变量，不用手动创建temp了
	fmt.Printf("s: %v", s)
	fmt.Printf("i: %v", i)
}

func TestSwap(t *testing.T) {
	a := 1
	b := 2
	c := 3
	d := 4
	a, b, c, d = b, c, a, d
	fmt.Println(a, b, c, d)
}

type Fragment interface {
	Exec() error
}

type GetPodAction struct {
}

func (g GetPodAction) Exec() error {
	return nil
}

func TestAssign(t *testing.T) {
	var fragment Fragment = new(GetPodAction)
	//var fragment3 Fragment = GetPodAction
	var fragment4 Fragment = &GetPodAction{}
	var fragment2 Fragment = GetPodAction{}
	fmt.Println(fragment)
	fmt.Println(fragment2)
	//fmt.Println(fragment3)
	fmt.Println(fragment4)
}

func TestEqual(t *testing.T) {
	//函数不可比较
	//var fn1 = func() {}
	//var fn2 = func() {}
	//if fn1 != fn2 {
	//	fmt.Println("equal")
	//}
}

type People struct {
	Name string `json:"name"` //首字母大写才能被json访问
}

func TestOutput(t *testing.T) {
	js := `{"name":"seekload"}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}

type foo2 struct {
	bar int
}

func TestRun(t *testing.T) {
	//var f foo2
	//f.bar, tmp := 1, 2
	//fmt.Println(tmp)
}

func TestSubBlock(t *testing.T) {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		i, x := 2, 2 //代码段局部变量作用范围
		fmt.Println(i, x)
	}
	fmt.Println(x)
}

func TestIncr(t *testing.T) {
	//data := []int{1, 2, 3}
	//i := 0
	//i++
	//fmt.Println(data[i++])
	//go不允许在这里使用++等
}

// const 连续赋值时可以胜率类型和值
const (
	x uint16 = 120
	y
	s = "abc"
	z
)

func TestConstDefault(t *testing.T) {
	fmt.Printf("%T %v", y, y)
	fmt.Printf("%T %v", z, z)
}

func TestCanCompiled(t *testing.T) {
	//变量赋值可以省略类型,类型推断
	const x = 123
	const y = 1.23
	const c = "hello"
	var z string = "hello"
	fmt.Println(z)
}

func TestIsError(t *testing.T) {
	//a := sync.Map{}
	//atomic.AddInt32()
	//atomic.LoadInt64()
	x := []int{1, 2}
	_ = x
}
