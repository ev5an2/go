package main

import "fmt"

func wanshu(a int) int {
	var sum int = 0
	for i:=1;i<a;i++{
		if a%i == 0 {
			sum+=i
			continue
		}
	}
	return sum
}

func main() {
	for i:=1;i<1001;i++{
		//因子之和
		if i == wanshu(i){
			fmt.Println(i)
		}

	}
}
