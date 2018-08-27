package main

import (
	"fmt"
)

//非阻塞通道
func main()  {

	messages :=make(chan string)
	signals :=make(chan bool)

	select {
	case msg :=<- messages:
		fmt.Println("received message :",msg)
	default:
		fmt.Println("no message send")
	}

	msg :="hi"
	select {
	case messages <-msg:
		fmt.Println("send message:",msg)
	default:
		fmt.Println("no message send")
	}

	select {
	case msg :=<-messages:
		fmt.Println("received message:",msg)
	case sig :=<-signals:
		fmt.Println("received signals ",sig)
	default:
		fmt.Println("no activity")
	}

}




