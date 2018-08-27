package main

import (
	"sort"
	"fmt"
)

//使用函数自定义排序
func main()  {

	fruits :=[]string{"peach","banana","kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println("fruits:",fruits)
}

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i,j int)  {
	s[i],s[j] =s[j],s[i]
}

func (s ByLength) Less(i,j int) bool  {
	return len(s[i])<len(s[j])
}
