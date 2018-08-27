package main

import "fmt"

//结构体
func main()  {
	fmt.Println(person{name:"Alice",age:30})

	fmt.Println(person{"Bob",20})

	fmt.Println(person{name:"pop"})

	fmt.Println(&person{name:"Ann",age:40})

	s:=person{name:"sean",age:50}

	fmt.Println(s.name)

	sp:=&s
	fmt.Println(sp.name)
	fmt.Println(sp.age)

	sp.age=51
	fmt.Println(sp.age)

}

type person struct{
	name string
	age int
}







