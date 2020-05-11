package main

//数组

import (
	"fmt"
)



func main() {
	var a [5]int
	a[0] = 100
	a[1] = 200
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	for i, v := range a {
		fmt.Printf("a[%d]%d\n", i, v)
	}

}
