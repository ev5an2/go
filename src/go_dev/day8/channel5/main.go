package main

import (
	"fmt"
)

func send(ch chan int,exitChan chan bool){
	for i:=0;i<10;i++{
		ch<-i
	}
	exitChan<-true
	close(ch)
}
func recv(ch chan int ,exitChan chan bool){
	for {
		v,ok:=<-ch
		if !ok{
			break
		}
		fmt.Println(v)
	}
	exitChan<-true
}
func main() {
	ch:=make(chan int,10)
	exitChan :=make(chan bool,2)
	go send(ch,exitChan)
	go recv(ch,exitChan)
	//如果后面不写，程序执行是没有结果，因为上面是两个协程，主进程退出了
	//可以用sleep10秒的方式保证协程执行完，也可以用exitChan

	for i:=0;i<2;i++{
		<-exitChan
	}

}
