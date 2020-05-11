package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}
func(a Student)Print(){
	fmt.Println(".....")
	fmt.Println(a)
}

func(a Student)Set(name string,age int){
	a.Name=name
	a.Age= age
}
func TestStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct{
		fmt.Println("expect struct")
		return
	}
	num := val.Elem().NumField()
	val.Elem().Field(0).SetString("STR001")

	for i:=0;i<num; i++{
		fmt.Printf("%d %v\n",i,val.Elem().Field(i))
	}
	fmt.Printf("struct %d fields\n", num)

	numMethod := val.Elem().NumMethod()
	fmt.Printf("struct %d Method\n", numMethod)

	var params []reflect.Value
	val.Elem().Method(0).Call(params)
}

func main() {
	var a Student = Student{
		Name: "stu1",
		Age:  30,
	}
	TestStruct(&a)
}
