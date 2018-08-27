package main

import "fmt"


//递归
func main()  {
	recurve:=recurve(7)
	fmt.Println(recurve)
}

func recurve(n int) int  {
	if n==0{
		return 1
	}
	return n*recurve(n-1)
}
