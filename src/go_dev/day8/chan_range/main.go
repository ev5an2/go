package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int ,1000)
	for i:=0; i<1000;i++{
		ch <-i
	}
	//一定要关掉，否则会报错
	close(ch)
	for v:=range ch{
		fmt.Println(v)
	}
}

