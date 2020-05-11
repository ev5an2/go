package main
//尾部插入
import (
	"fmt"
	"math/rand"
)

type Student struct {
	Name string
	Age int
	Score float32
	next *Student


}

func trans(p *Student){
	for p!=nil {
		fmt.Println(*p)
		p = p.next
	}
}

//func main(){
//	var head Student
//	head.Age = 18
//	head.Name = "tang"
//	head.Score = 80
//
//	var stu1 Student
//	stu1.Age = 1
//	stu1.Name = "stu1"
//	stu1.Score = 1
//	//stu1.next = nil
//
//	head.next = &stu1
//
//	//var p *Student = &head
//	//trans(&head)
//
//	var stu2 Student
//	stu2.Name = "stu2"
//	stu2.Age = 23
//	stu2.Score = 23
//	stu1.next = &stu2
//	trans(&head)
//
//	var stu3 Student
//	stu3.Name = "stu3"
//	stu3.Age = 233
//	stu3.Score = 233
//	stu2.next = &stu3
//	trans(&head)
//}

func main(){
	var head Student
	head.Age = 18
	head.Name = "tang"
	head.Score = 80

	var tail = &head
	for i:=0;i<20;i++{
		var stu = Student{
			Name:  fmt.Sprintf("stu%d",i),
			Age:   rand.Intn(100),
			Score: rand.Float32()*100,
		}
		tail.next = &stu
		tail = &stu
	}
	trans(&head)
}