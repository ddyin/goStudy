package main

import (
	"os"
	"fmt"
)

//命令行参数
func main()  {

	argsWtihProg := os.Args
	argsWtihoutProg := os.Args[1:]

	arg := os.Args[0]

	fmt.Println(argsWtihoutProg)
	fmt.Println(argsWtihProg)
	fmt.Println(arg)

}


