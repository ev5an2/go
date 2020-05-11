package main

import (
	"fmt"
	"go_dev/day6/lianxi2/balance"
	"math/rand"
	"time"
)

func main() {
	var insts []*balance.Instance
	for i := 0; i < 16; i++ {
		inst := &balance.Instance{
			Host: fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255)),
			Port: 8080,
		}
		insts = append(insts, inst)
	}
	var a balance.Random
	for {
		res, err := a.DoBalance(insts)
		if err != nil {
			println(err)
			continue
		}
		time.Sleep(time.Second)
		fmt.Println(res)
	}

}
