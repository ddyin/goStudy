package main

import "os"

//panic(出乎意料的错误发生)
func main()  {
	panic("a problem")

	_,err := os.Create("/tmp/file")
	if err!=nil {
		panic(err)
	}
}



