package main

import (
	"os"
	"fmt"
	"strings"
)

//环境变量
func main()  {

	os.Setenv("FOO","1")
	fmt.Println("FOO:",os.Getenv("FOO"))
	fmt.Println("BAR:",os.Getenv("BAR"))

	for _,e := range os.Environ(){
		pair := strings.Split(e,"=")
		fmt.Println(pair[0])
	}
}
