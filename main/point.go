package main

import "fmt"

//指针
func main()  {

	i:=1
	fmt.Println("initial:",i)

	zeroval(i)

	fmt.Println("zeroval:",i)

	zeroptr(&i)
	fmt.Println("zeroptr:",i)

	fmt.Println("pointer:",&i)

}

func zeroval(ival int)  {
	ival=0
}

func zeroptr( iptr *int)  {
	*iptr=0
}


