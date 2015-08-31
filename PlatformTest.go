package main

import (
	"fmt"
	"runtime"
)

func PlatformTest() {
	fmt.Println("区分不同平台测试")
	fmt.Println(runtime.GOOS)
}
