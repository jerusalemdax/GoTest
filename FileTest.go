package main

import (
	"fmt"
	"io/ioutil"
)

func FileTest() {
	fmt.Println("文件操作测试")
	ReadFile()
}

func ReadFile() {
	buffer, err := ioutil.ReadFile("servers.xml")
	if err != nil {
		fmt.Println("open file error: ", err)
		return
	}
	fmt.Println(string(buffer))
}
