package balance

import (
	"errors"
	"math/rand"
)

type Random struct {
}

func (r *Random) DoBalance(insts []*Instance) (inst *Instance, err error) {
	lens:=len(insts)

	if lens ==0 {
		err = errors.New("no instance")
		return
	}

	inst =insts[rand.Intn(lens)]

	return

}
