package main
//头部插入
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
	//fmt.Println(*p)
	for p!=nil {
		fmt.Println(*p)
		p = p.next
	}
}

//func mainErr(){
//	var head Student
//	head.Age = 18
//	head.Name = "tang"
//	head.Score = 80
//
//	for i:=0;i<10;i++{
//		stu := Student{
//			Name:  fmt.Sprintf("stu%d",i),
//			Age:   rand.Intn(100),
//			Score: rand.Float32()*100,
//		}
//		stu.next = &head
//		//head = stu---->{head.Name = stu.Name,head.Age= stu.Age,head.next=stu.next}
//		//错误在于head.next = stu.next = &head.也就是把head的地址绑定到了head的next上。在下次循环中把head的地址给stu.next.
//		//相当于stu.next = head.next 相当于stu1.next = stu0.next.所以地址全部相同
//		head = stu
//		fmt.Println(stu)
//		time.Sleep(time.Second)
//	}
//	trans(&head)
//}

func insertHead(p **Student){
	for i:=0;i<10;i++{
		stu := Student{
			Name:  fmt.Sprintf("stu%d",i),
			Age:   rand.Intn(100),
			Score: rand.Float32()*100,
		}
		stu.next = *p
		*p = &stu
	}
}
func delNode(p *Student){
	var prev *Student = p
	for p!=nil{
		if p.Name == "stu6"{
			prev.next = p.next
			break
		}
		prev = p
		p =p.next
	}
}

func addNode(p *Student,newNode *Student){
	for p!=nil{
		if p.Name == "stu5"{
			newNode.next = p.next
			p.next = newNode
			break
		}
		p =p.next
	}
}


func main(){
	var head *Student = &Student{}
	//var head *Student = new(Student)
	head.Age = 18
	head.Name = "tang"
	head.Score = 80

	insertHead(&head)
	trans(head)
	//删除某个节点
	delNode(head)
	trans(head)

	var newNode *Student = new(Student)
	newNode.Age = 18
	newNode.Name = "LIU"
	newNode.Score = 100
	//插入新节点
	addNode(head,newNode)
	trans(head)
}

func main2(){
	var head *Student = &Student{}
	//var head *Student = new(Student)
	head.Age = 18
	head.Name = "tang"
	head.Score = 80

	for i:=0;i<10;i++{
		stu := Student{
			Name:  fmt.Sprintf("stu%d",i),
			Age:   rand.Intn(100),
			Score: rand.Float32()*100,
		}
		stu.next = head
		head = &stu
	}
	trans(head)
}