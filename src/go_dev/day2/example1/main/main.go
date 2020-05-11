package main

import "fmt"

func list(n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d+%d=%d\n", i, n-i, n)
	}
}

func main(){
	list(10)
}

//func main() {
//	n := 100
//	for i := 0; i < n; i++ {
//		fmt.Println(i, "+", n-i, "=", n)
//	}
//
//}
