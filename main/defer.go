package main

import (
	"os"
	"fmt"
)

//确保程序在执行结束前执行，相当于java中的finally
func main()  {

	f :=createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File  {
	fmt.Println("creating")
	f,err :=os.Create(p)

	if err!=nil{
		panic(err)
	}
	return f
}

func writeFile(f *os.File)  {
	fmt.Println("writing...")
	fmt.Fprintln(f,"data")
}

func closeFile(f *os.File)  {
	fmt.Println("closing")
	f.Close()
}




