package balance

import (
	"errors"
	"math/rand"
)

type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(list []*Instance) (ins *Instance, err error) {

	size := len(list)
	if size == 0 {
		err = errors.New(" Instance size is zero ")
		return
	}
	ins = list[rand.Intn(size)]
	return
}
