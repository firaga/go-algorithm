package josephRing

import (
	"fmt"
	"math/rand"
	"time"
)

//.编写一个 Go 程序，30分钟内完成：模拟一个简单的生产者-消费者问题。
//有一个生产者 goroutine 负责生成随机整数，并将它们放入一个有限大小的缓冲区（buffer）中，
//同时有多个消费者 goroutine 从缓冲区中取出整数并进行处理。
func q1() {
	var triangle [][]int
	dp := make([][]int, len(triangle))
	for i := range dp {
		dp[i] = make([]int, len(triangle[0]))
	}
	numJobs := 5
	numWorks := 2
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	for i := 0; i < numWorks; i++ {
		go consumer(i, jobs, results)
	}
	producer(jobs, numJobs)
	close(jobs)
	for a := 0; a < numJobs; a++ {
		f := <-results
		fmt.Println(f)
	}
}

func producer(jobs chan<- int, numJobs int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numJobs; i++ {
		randNum := rand.Intn(100)
		jobs <- randNum
	}
}

func consumer(seq int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("consumer", seq, "handle job numer", j)
		time.Sleep(time.Second)
		results <- j
	}
}
