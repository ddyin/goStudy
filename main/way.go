package main

import "fmt"

//方法
func main()  {

	r:=rect{width:10,height:5}
	fmt.Println("area:",r.area())
	fmt.Println("preim:",r.perim())

	rp:=&r
	fmt.Println("area:",rp.area())
	fmt.Println("preim:",rp.perim())

}

type rect struct {
	width,height int
}

func (r *rect)area() int  {
	return r.width*r.height
}

func (r * rect)perim() int  {
	return 2*r.width+2*r.height
}
