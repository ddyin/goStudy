package main

import "fmt"

//多返回值
func main()  {

	res1,res2:=cal()
	fmt.Println("res1:",res1,",res2:",res2)

	_,d:=cal()
	fmt.Println(d)
}

func cal() (int,int)  {
	return 11,33
}
