package main

import "fmt"

func testSlice(){
	var slice []int
	var arr [5]int = [...]int{1,2,3,4,5}
	slice = arr[4:5]
	slice = arr[:]
	slice = arr[1:]
	slice = arr[:len(slice)]
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	slice = arr[2:3]
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}




func main() {
	testSlice()
}

