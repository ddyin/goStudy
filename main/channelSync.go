package main

import (
	"fmt"
	"time"
)

//通道同步
func main()  {

	done := make(chan bool ,1)
	go worker(done)

	<-done
}

func worker(done chan bool)  {

	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

