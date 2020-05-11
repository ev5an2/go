package main

import (
	"fmt"
)

func reverse(str string) string {
	var result string
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		result = result + fmt.Sprintf("%c", str[strlen-i-1])
	}
	return result
}

func reverse1(str string) string {
	var result []byte
	tmp := []byte(str)
	length := len(str)
	for i := 0; i < length; i++ {
		result = append(result, tmp[length-i-1])
	}
	return string(result)
}

func main() {
	var str1 = "hello"
	var str2 = "world"
	str3 := fmt.Sprintf("%s %s", str1, str2)
	//n := len(str3)
	//fmt.Println(n)
	n := reverse(str3)
	fmt.Println(n)
}
