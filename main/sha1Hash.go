package main

import (
	"crypto/sha1"
	"fmt"
)

//SHA1散列(哈希)
func main()  {
	s := "sha1 this string "
	h := sha1.New()
	h.Write([]byte(s))

	//fmt.Println(h)
	//fmt.Println(h.Size())

	bs := h.Sum(nil)
	fmt.Println(s)

	fmt.Printf("%x\n",bs)
}
