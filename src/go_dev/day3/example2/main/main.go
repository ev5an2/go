package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "   hello world abc  \n"
	result := strings.Replace(str, "world", "you", 1)
	fmt.Println(result)

	count := strings.Count(str, "l")
	fmt.Println(count)

	result = strings.Repeat(str, 3)
	fmt.Println(result)

	result = strings.ToLower(str)
	fmt.Println(result)

	result = strings.ToUpper(str)
	fmt.Println(result)

	result = strings.TrimSpace(str)
	fmt.Println("trimSpace:", result)

	result = strings.Trim(str, " \n\r")
	fmt.Println("trim:", result)

	result = strings.TrimLeft(str, " \n\r")
	fmt.Println("trimLeft:", result)

	result = strings.TrimRight(str, " \n\r")
	fmt.Println("trimRight:", result)

	splitResult := strings.Fields(str)
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}
	splitResult = strings.Split(str,"l")
	for i := 0; i < len(splitResult); i++ {
		fmt.Println(splitResult[i])
	}

	str2 := strings.Join(splitResult, "l")
	fmt.Println("join:", str2)

	str2 = strconv.Itoa(1000)
	fmt.Println("itoa:", str2)

	number, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("can not convert to int", err)
	}
	fmt.Println("number:", number)

}
