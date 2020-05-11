package main

import (
	"errors"
	"fmt"
)

func initConfig() (err error)  {
	return errors.New("init config failed")
}
func test(){
	//defer func(){
	//	if err:= recover();err!=nil{
	//		fmt.Println(err)
	//	}
	//}()

	err:=initConfig()
	if err !=nil{
		panic(err)
	}
	b:=0
	a:=100/b
	fmt.Println(a)
	return
}
func main(){
	//for {
		test()
		//time.Sleep(time.Second)
	//}


	var a []int
	a = append(a,10,20)
	a = append(a,a...)
	fmt.Println(a)




	var i int
	fmt.Println(i)
	j:=new(int)
	*j = 100
	fmt.Println(*j)
}
