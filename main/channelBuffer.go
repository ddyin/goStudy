package main

import "fmt"

//通道缓冲
func main()  {

	messages := make(chan string,2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

}




