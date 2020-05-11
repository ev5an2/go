package main

func main() {
	//var intChan chan int
	//intChan = make(chan int,10)
	//intChan <- 10
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string,10)
	m:=make(map[string]string,16)
	m["stu01"] = "001"
	m["stu02"] = "002"

	mapChan<-m

}
