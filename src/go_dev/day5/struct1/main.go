package main

import "fmt"

type Student struct {
	Name string
	Age int
	score float32


}


func main(){
	var stu Student
	stu.Age = 18
	stu.Name = "tang"
	stu.score = 80.0

	//var stu1 *Student = &Student{
	//	Age:20,
	//	Name:"jia",
	//}
	//fmt.Println(stu1.Name)

	//var stu3 = Student{
	//	Name:  "12",
	//	Age:   0,
	//	score: 0,
	//}
	fmt.Println(stu3)
	fmt.Printf("Name:%p\n",&stu.Name)
	fmt.Printf("Age:%p\n",&stu.Age)
	fmt.Printf("score:%p\n",&stu.score)
}