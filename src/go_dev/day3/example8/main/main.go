package main

import "fmt"

func test(a,b int) int {
	result:=func(a1 int,b1 int)int{
		return a1+b1
	}(a,b)
	return result
}
func main() {

	fmt.Println(test(100,200))
}
