package main

import (
	"fmt"
	"time"
)

func main()  {

	i:=2
	fmt.Printf("write",i,"as")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday,time.Sunday:
		fmt.Println("weekend")
	case time.Monday:
		fmt.Println("周一")
	default:
		fmt.Println("这是除周一外的工作日")

	t:=time.Now()
	fmt.Println(t)

	switch  {
	case t.Hour() <12:
		fmt.Println("这是上午")
	default:
		fmt.Println("下午来了")
	}
		
		


		
		
	
	}








}
