package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

	fmt.Println("执行命令测试用例")
	fmt.Println("执行ls -al命令")
	cmd := exec.Command("ls", "-al")
	out, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "执行命令错误: %s ", err)
		return
	}
	fmt.Fprintf(os.Stdout, "执行结果: %s", out)

	fmt.Println("区分不同平台测试")
	fmt.Println(runtime.GOOS)

	fmt.Println("protobuf测试")
}
