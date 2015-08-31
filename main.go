package main

import (
	"GoTest/example"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
	"os/exec"
	"runtime"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

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
	test := &example.Test{
		Label: proto.String("this is protobuf test"),
		Type:  proto.Int32(17),
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		fmt.Println("marshaling error: ", err)
	}
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		fmt.Println("unmarshaling error: ", err)
	}
	if test.GetLabel() != newTest.GetLabel() {
		fmt.Println("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
	fmt.Println(newTest.GetLabel())
	fmt.Println(newTest.GetType())
	fmt.Println(newTest.GetOptionalgroup())

	fmt.Println("解析json测试")
	var s Serverslice
	str2 := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str2), &s)
	fmt.Println(s)

	fmt.Println("生成json测试")
	var s2 Serverslice
	s2.Servers = append(s2.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s2.Servers = append(s2.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s2)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
