package main

import (
	"flag"
	"fmt"
)

func main(){
	var confPath string
	var logLevel int
	//第三个参数是默认值
	//go run main.go -c 123 -d 456
	flag.StringVar(&confPath,"c","","please input conf path")
	flag.IntVar(&logLevel,"d",10,"please input log level")

	flag.Parse()
	fmt.Println("path",confPath)
	fmt.Println("log level",logLevel)
}

