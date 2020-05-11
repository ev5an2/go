package main

import "fmt"

func add(a int, arg ...int) int {
	var sum int = a
	for i := 0; i < len(arg); i++ {
		sum += arg[i]
	}
	return sum
}

func spell(c string, arg...string) (result string) {
	result  = c
	for i := 0; i < len(arg); i++ {
		result += arg[i]
	}
	return
}
func main() {
	a := add(3, 1, 2)
	fmt.Println(a)

	b := spell("tang", "jie", "hello", "!")
	fmt.Println(b)
}
