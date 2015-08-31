package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type Serverslice struct {
	Servers []Server
}

func JsonTest() {
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
