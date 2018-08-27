package main

import (
	"fmt"
	"time"
)

//工作池
func main()  {

	jobs := make(chan int,100)

	results := make(chan int,100)

	for w := 1;w <=3;w++ {
		go workers(w,jobs,results)
	}

	// 这里我们发送 9 个 `jobs`，然后 `close` 这些通道
	// 来表示这些就是所有的任务了。
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	// 最后，我们收集所有这些任务的返回值。
	for a := 1; a <= 9; a++ {
		<-results
	}
}

func workers(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

