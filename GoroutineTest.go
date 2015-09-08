package main

import (
	"fmt"
	"time"
)

var c chan int

func GoroutineTest() {
	fmt.Println("这是Goroutine测试")
	c = make(chan int)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	<-c
	<-c
}

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready")
	c <- 1
}
