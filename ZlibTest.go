package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io/ioutil"
)

func ZlibTest() {
	fmt.Println("Zlib压缩测试")
	CompressTest()
}

func CompressTest() {
	buffer, err := ioutil.ReadFile("servers.xml")
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	fmt.Println("压缩前大小:", len(buffer))
	var b bytes.Buffer
	//w := zlib.NewWriter(&b)
	w, err := zlib.NewWriterLevel(&b, zlib.BestSpeed)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	w.Write(buffer)
	w.Close()
	fmt.Println("压缩后大小:", b.Len())
	fmt.Println("压缩结果")
	fmt.Println(b.Bytes())
}
