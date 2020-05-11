package main

import "fmt"

func add(a int, b int, c chan int) {
	var sum int
	sum = a + b
	c <- sum

}
func main() {
	//var pipe chan int
	//pipe = make(chan int, 1)
	//go add(2, 5, pipe)
	//
	//sum := <-pipe
	//fmt.Println("sum =", sum)

	//var c int
	//c = add(100,200)
	//go test_goroute(200,200)
	//fmt.Println("add(100,200)=",c)
	//for i := 0; i < 100; i++ {
	//	go test_print(i)
	//}
	//time.Sleep(time.Second)
	//test_pipe()

	sum, avg := cale(100, 200)
	fmt.Println("sum=", sum, "avg=", avg)
}
