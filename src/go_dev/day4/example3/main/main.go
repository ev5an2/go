package main

import "fmt"

func factor(n int) int {
	if n <= 1 {
		return 1
	}
	return factor(n-1)+factor(n-2)
}
func main() {
	n:=5
	for i:=0;i<n;i++{
		fmt.Printf("%d ",factor(i))
	}
	//fmt.Println(factor(5))

}
