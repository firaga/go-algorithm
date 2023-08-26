package main

import "fmt"

func main() {
	var ch chan int
	v, ok := <-ch
	fmt.Println(v, ok)
}
