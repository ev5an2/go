package main

import (
	"fmt"

)

func isRight(a string) bool {
	var str string
	for i:=0;i<len(a);i++{
		str+=string(a[len(a)-i-1])
	}
	fmt.Println(str)
	if str == a {
		return true
	}else{
		return false
	}

}

func main() {
	var a string
	fmt.Scanf("%s",&a)
	b:=isRight(a)
	fmt.Println(b)
}
