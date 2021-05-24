package main

import (
	"fmt"
	"time"
)

func main() {
	ch := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; ; i++ {
				ch <- i
			}
		}()
		return ch
	}()
	for v := range ch {
		fmt.Println(v)
		if v == 5 {
			break
		}
	}
	time.Sleep(60 * time.Second)
}
