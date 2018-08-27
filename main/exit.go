package main

import (
	"fmt"
	"os"
)

//退出
func main()  {

	defer fmt.Println("!")

	os.Exit(3)
}

