package main

import (
	"fmt"
	"os"
	"os/exec"
)

func CommandTest() {
	fmt.Println("执行命令测试用例")
	fmt.Println("执行ls -al命令")
	cmd := exec.Command("ls", "-al")
	out, err := cmd.Output()
	if err != nil {
		fmt.Fprintf(os.Stderr, "执行命令错误: %s ", err)
		return
	}
	fmt.Fprintf(os.Stdout, "执行结果: %s", out)
}
