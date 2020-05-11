package main

import "fmt"

func swap(a *int, b *int) {
	c := *a
	*a = *b
	*b = c
	return
}

func main() {
	a := 3
	b := 4
	swap(&a, &b)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
}

//func change(a int,b int) (a1 int,b1 int){

//	return b,a
//}
//
//
//func main() {
//	a :=3
//	b :=4
//	a,b = change(a,b)
//	fmt.Println("a=",a,"\n b=",b)
//}
