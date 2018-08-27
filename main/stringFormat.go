package main

import (
	"fmt"
	"os"
)

//字符串格式化
func main()  {
	p:=point{1,2}

	fmt.Println("%v\n",p)

	fmt.Println("%+v\n",p)

	fmt.Println("%#v\n",p)

	fmt.Println("%T\n",p)

	fmt.Println("%t\n",true)

	fmt.Println("%d\n",133)

	fmt.Println("%b\n",14)

	fmt.Println("%c\n",33)

	fmt.Println("%x\n",456)

	fmt.Println("%f\n",78.9)

	fmt.Printf("%e\n", 123400000.0)

	fmt.Printf("%E\n", 123400000.0)

	fmt.Printf("%s\n", "\"string\"")

	fmt.Printf("%q\n", "\"string\"")

	fmt.Printf("%x\n", "hex this")

	fmt.Printf("%p\n", &p)

	fmt.Printf("|%6d|%6d|\n", 12, 345)

	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("|%6s|%6s|\n", "foo", "b")

	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)

	fmt.Fprintln(os.Stderr,"an %s\n","error")

}

type point struct {
	x,y int
}


