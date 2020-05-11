package main

import (
	"fmt"
	"math"
)

func isRight(a int) bool {
	if (a <= 99 || a >= 1000) {
		return false
	} else {
		a1 := int(math.Floor(float64(a/100)))
		a2 := int(math.Floor(float64((a - a1*100)/10)))
		a3 := int(math.Floor(float64(a - a1*100 - a2*10)))
		//fmt.Println(a1,a2,a3)
		if (a1*a1*a1+a2*a2*a2+a3*a3*a3 == a) {
			return true
		} else {
			return false
		}

	}

}

func main() {
	//isRight(125)
	for j := 100; j < 1000; j++ {
		if (isRight(j)) {
			fmt.Println(j)
		}
	}
}
