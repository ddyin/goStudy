package main

import (
	"sort"
	"fmt"
)

//排序
func main()  {

	strs:=[]string{"c","a","b"}
	sort.Strings(strs)
	fmt.Println("strings:",strs)

	ints :=[]int{7,2,4}
	sort.Ints(ints)
	fmt.Println("Ints:",ints)

	//是否已排序
	s:=sort.IntsAreSorted(ints)
	fmt.Println("sorted:",s)

}
