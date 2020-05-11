package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	UserName string
	NickName string
	Age      int
	Birthday string
	Sex      string
	Email    string
	Phone    string
}
func testStruct(){
	user1 := &User{
		UserName: "user01",
		NickName: "tt",
		Age:      18,
		Birthday: "2019/1/2",
		Sex:      "男",
		Email:    "54533148@qq.com",
		Phone:    "111",
	}

	data,err :=json.Marshal(user1)
	if err !=nil{
		fmt.Printf("json marshal failed,err:",err)
		return
	}

	fmt.Printf("%s\n",string(data))
	fmt.Printf("%T\n",data)
}
func testInt(){
	var age = 100
	data,err :=json.Marshal(age)
	if err !=nil{
		fmt.Printf("json marshal failed,err:",err)
		return
	}

	fmt.Printf("%s\n",string(data))
}

func testMap(){
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"

	data,err :=json.Marshal(m)
	if err !=nil{
		fmt.Printf("json marshal failed,err:",err)
		return
	}

	fmt.Printf("%s\n",string(data))
}

func testSlice(){
	var s []map[string]interface{}
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"

	s =append(s,m)
	m = make(map[string]interface{})
	m["username"] = "user2"
	m["age"] = 28
	m["sex"] = "female"
	s =append(s,m)

	data,err :=json.Marshal(s)
	if err !=nil{
		fmt.Printf("json marshal failed,err:",err)
		return
	}

	fmt.Printf("%s\n",string(data))
}

func main(){
	//testStruct()
	//testMap()
	testSlice()
}