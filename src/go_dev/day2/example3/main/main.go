package main

import (
	"fmt"
	"time"
)

const (
	Man    = 1
	Female = 2
)

func main() {
	second := time.Now().Unix()
	if (second % Female == 0) {
		fmt.Println("famale")
	} else {
		fmt.Println("man")
	}
}
