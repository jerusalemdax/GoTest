package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World")
	arg_num := len(os.Args)
	fmt.Println("命令行参数为:")
	for i := 0; i < arg_num; i++ {
		fmt.Println(os.Args[i])
	}

	fmt.Println("命令行解析库flag测试用例")
	var str = flag.String("t", "test2", "test string argument")
	var num = flag.Int("i", 1234, "test num argument")
	var bo = flag.Bool("b", false, "test bool argument")
	flag.Parse()
	fmt.Println("t参数为: %s", *str)
	fmt.Println("n参数为: %d", *num)
	fmt.Println("b参数为: %t", *bo)
}
