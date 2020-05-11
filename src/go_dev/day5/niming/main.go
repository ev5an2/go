package main

//匿名字段
import (
	"fmt"
	"time"
)

type Car struct {
	name string
	age int
}
type Train struct{
	Car
	int
	start time.Time
	age int
}
func main() {
	var t Train
	t.name = "train"
	t.age = 100

	t.Car.name = "train1"
	t.Car.age = 200

	t.int = 200
	fmt.Println(t)
}
