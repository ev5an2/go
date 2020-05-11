package main

import "fmt"

type Test interface {
}

type Carer interface {
	GetName() string
	Run()
	DiDi()
}
type BYD struct {
	Name string
}
type BMW struct {
	Name string
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run()  {
	fmt.Printf("%s is running\n",p.Name)
}

func (p *BMW) DiDi()  {
	fmt.Printf("%s is didi\n",p.Name)
}

func (p *BYD) GetName() string {
	return p.Name
}

func (p *BYD) Run()  {
	fmt.Printf("%s is running\n",p.Name)
}

func (p *BYD) DiDi()  {
	fmt.Printf("%s is didi\n",p.Name)
}

func main() {
	var car Carer
	var bmw BMW
	bmw.Name = "bmw"
	//这里加地址符是因为用的指针*类型实现的接口，可以把*和&都去掉
	car = &bmw
	car.Run()

	byd:=&BYD{
		Name:"BYD",
	}
	car = byd
	car.Run()
}















func main1() {
	var a interface{}
	var b int
	//b实现了a接口，b给a赋值
	a = b
	fmt.Printf("type of %T\n", a)
}
