package main

import (
	"time"
	"fmt"
)

//定时器
func main()  {

	timer1 := time.NewTimer(time.Second*2)

	<-timer1.C
	//time1.C: 0xc042084000
	fmt.Println("time1.C:",timer1.C)
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second	)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("Timer 2 stopped ")
	}
}



