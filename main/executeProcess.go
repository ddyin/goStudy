package main

import (
	"os/exec"
	"os"
	"syscall"
)

//执行进程
func main()  {


	binary,lookErr := exec.LookPath("ls")
	if lookErr !=nil {
		panic(lookErr)
	}

	args := []string{"ls","-a","-1","-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary,args,env)

	if execErr !=nil{
		panic(execErr)
	}

}



