package main

import (
	b64 "encoding/base64"
	"fmt"
)

//base64编码
func main()  {
	data := "abc123!?$*&()'-=@~"

	//URL的编码不同的是编码后添加了+ 或者-的区别，但解码出来数据都相同

	//编码
	sEnc := b64.StdEncoding.EncodeToString([]byte (data))
	fmt.Println(sEnc)
	//解码
	sDec,_ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println(string(sDec))
	//编码
	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	//解码
	uDec,_ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println(string(uDec))

}


