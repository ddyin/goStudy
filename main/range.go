package main

import "fmt"

//遍历
func main()  {

	nums:=[]int{2,3,4}
	fmt.Println("nums:",nums)

	sum :=0
	for _,num:=range nums{
		sum+=num
	}

	fmt.Println("sum:",sum)

	for i,num:=range nums{
		if num == 3 {
			fmt.Println("index:",i)
		}
	}

	kvs:=map[string]string{"a":"apple","b":"banana"}
	for k,v:=range kvs{
		fmt.Println("%s -> %s\n",k,v)
	}

	for i,c:=range "go"{
		fmt.Println(i,c)
	}
}










