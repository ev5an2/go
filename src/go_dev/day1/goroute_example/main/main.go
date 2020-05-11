package main

import (
	"fmt"
	"go_dev/day1/goroute_example/goroute"
)

func main() {

	pipe := make(chan int, 1)
	goroute.Add(1, 1, pipe)
	sum:=<-pipe
	//打印整型二进制
	//fmt.Printf("%b",sum)
	//打印十进制
	//fmt.Printf("%d",sum)
	//打印十六进制
	//fmt.Printf("%x",sum)
	//浮点数
	//fmt.Printf("%f",1.2)
	//fmt.Printf(%s,sum)
}
