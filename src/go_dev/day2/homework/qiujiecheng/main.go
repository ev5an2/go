package main

import "fmt"

func sum(n int) uint64 {
	var s uint64 = 1
	var sum uint64 = 0
	for i := 1; i <= n; i++ {
		s = s * uint64(i)
		sum += s
	}
	return sum
}

func jiechen(a int) int {
	p := 1
	for i := 1; i <= a; i++ {
		p = p * i
	}
	return p
}

func main() {
	n := 3
	sum := 0
	for j := 1; j < n+1; j++ {
		sum = sum + jiechen(j)
	}
	fmt.Println(sum)
}
