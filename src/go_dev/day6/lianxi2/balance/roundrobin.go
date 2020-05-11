package balance

import (
	"errors"
)

type RoundRobin struct {
	curIndex int
}

func (r RoundRobin) DoBalance(insts []*Instance) (inst *Instance, err error) {
	lens := len(insts)
	if lens == 0 {
		err = errors.New("no instance")
		return
	}

	if r.curIndex > lens {
		r.curIndex = 0
	}
	inst = insts[r.curIndex]
	r.curIndex = (r.curIndex + 1) % lens
	return

}
