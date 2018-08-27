package main

import "fmt"

//通道
func main()  {
	messages :=make(chan string)

	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)
}

