package main

import (
	"fmt"
)

func tongji(a string) (int, int, int, int) {
	var b, c, d, e int
	for i := 0; i < len(a); i++ {
		switch {
		case 64 < a[i] && a[i] < 91:
			b += 1
		case 96 < a[i] && a[i] < 123:
			b += 1
		case 47 < a[i] && a[i] < 58:
			c += 1
		case a[i] == 32:
			d += 1
		default:
			e += 1
		}
	}
	return b, c, d, e

}

func main() {
	var a string
	fmt.Scanln("%s", &a)
	b, c, d, e := tongji(a)
	fmt.Println(b, c, d, e)
}
