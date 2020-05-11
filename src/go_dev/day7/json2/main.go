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

func testStruct() (ret string,err error){
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
		err = fmt.Errorf("json marshal failed,err:",err)
		return
	}
	ret =  string(data)
	return

}
func testMap() (ret string,err error){
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["username"] = "user1"
	m["age"] = 18
	m["sex"] = "man"

	data,err :=json.Marshal(m)
	if err !=nil{
		err = fmt.Errorf("json marshal failed,err:",err)
		return
	}

	ret = string(data)
	return
}

func test(){
	data,err :=testMap()
	if err != nil {
		fmt.Println("test struct failed,",err)
		return
	}
	var user1 User
	err = json.Unmarshal([]byte(data),&user1)
	if err != nil{
		fmt.Println("Unmarshal failed,",err)
		return
	}
	fmt.Println(user1)
	fmt.Printf("%T\n",user1)

}

func test2(){
	data,err :=testStruct()
	if err != nil {
		fmt.Println("test struct failed,",err)
		return
	}
	var m map[string]interface{}
	err = json.Unmarshal([]byte(data),&m)
	if err != nil{
		fmt.Println("Unmarshal failed,",err)
		return
	}
	fmt.Println(m)

}

func main(){
	test()
	test2()
}