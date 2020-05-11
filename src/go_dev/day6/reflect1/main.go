package main

import (
	"fmt"
	"reflect"
)
type Student struct {
	Name string
	Age  int

}
func test(b interface{}) {
	t := reflect.TypeOf(b)
	fmt.Println(t)
	v:=reflect.ValueOf(b)
	fmt.Println(v)
	k:=v.Kind()
	fmt.Println(k)
	g:=v.Interface()
	stu,ok :=g.(Student)
	if ok {
		fmt.Printf("%v %T\n",stu,stu)
	}
}

func testInt(b interface{}){
	val:=reflect.ValueOf(b)
	//Elem类似于在前面加*
	val.Elem().SetInt(100)
	c:=val.Elem().Int()
	fmt.Printf("get value  interface{} %d\n",c)
	fmt.Printf("%d\n",val.Elem().Int())
}
func main() {
	//var a int = 200
	var a Student=Student{
		Name: "stu1",
		Age:  10,
	}
	var b int = 1
	test(a)
	testInt(&b)
}
