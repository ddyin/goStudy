package main

import "fmt"
import (
	s "strings"
)

var p = fmt.Println

//字符串函数
func main()  {
	p("Contains:  ",s.Contains("test","es"))

	p("Count :",s.Count("testsssss","s"))

	p("HasPrefix :",s.HasPrefix("testsssss","te"))

	p("HasSuffix :",s.HasSuffix("testsssss","sss"))
	p("Index :",s.Index("testsssss","sss"))

	p("Join :",s.Join([]string{"a","b"},"---"))

	p("Repeat :",s.Repeat("a",5))

	p("Replace :",s.Replace("foo","o","0",-1))

	p("Replace :",s.Replace("foo","o","0",1))

	p("Split :",s.Split("a-b-c-d","-"))

	p("tolower :",s.ToLower("AAAA"))

	p("toupper :",s.ToUpper("aaaa"))

	p()
	p("length :",len("hahhahha "))

	p("char :","hello "[2])

	p("test:",s.Compare("j","k"))
}



