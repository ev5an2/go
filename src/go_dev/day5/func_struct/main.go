package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func (p *Student) init(name string, age int, score int) {
	p.Name = name
	p.Age = age
	p.Score = score
	fmt.Println(p)
}

func (p Student) get() Student{
	return p
}


func main() {
	var stu Student
	stu.init("stu",10,80)
	stu1:=stu.get()
	fmt.Println(stu1)
}
