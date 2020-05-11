package main

import (
	"fmt"
	"go_dev/day6/lianxi/balance"
	"hash/crc32"
	"math/rand"
)

type HashBalance struct {
}

func init(){
	balance.RegisterBalancer("hash",&HashBalance{})
}

func (p *HashBalance) DoBalance(insts []*balance.Instance, key ...string) (inst *balance.Instance, err error) {
	var defKey string = fmt.Sprintf("%d",rand.Int())
	if len(key) > 0 {
		defKey = key[0]
	}

	lens := len(insts)
	if lens == 0 {
		err = fmt.Errorf("no instance")
		return
	}

	crcTable :=crc32.MakeTable(crc32.IEEE)
	hashVal := crc32.Checksum([]byte(defKey),crcTable)
	println(int(hashVal))
	index := int(hashVal) % lens
	println(index)
	inst = insts[index]
	return
}
