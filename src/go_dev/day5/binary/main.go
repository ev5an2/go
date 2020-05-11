package main

import "fmt"

type Student struct {
	Name string
	Age int
	Score float32
	left *Student
	right *Student


}


func trans(root *Student){
	if root == nil {
		return
	}
	fmt.Println(root)
	trans(root.left)
	trans(root.right)
}

func main(){
	var root *Student = &Student{}
	//var head *Student = new(Student)
	root.Age = 18
	root.Name = "stu01"
	root.Score = 80


	var left1 *Student = &Student{}
	left1.Age = 18
	left1.Name = "stu02"
	left1.Score = 80
	root.left =left1

	var right1 *Student = &Student{}
	right1.Age = 18
	right1.Name = "stu04"
	right1.Score = 80
	root.right =right1

	var left2 *Student = &Student{}
	left2.Age = 18
	left2.Name = "stu03"
	left2.Score = 80
	left1.left =left2
	trans(root)
}

