package main

import "fmt"

//通道遍历
func main()  {

	queue :=make(chan string,2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue{
		fmt.Println("element :",elem)
	}
}

