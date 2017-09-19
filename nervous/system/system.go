package system

import (
	"fmt"
)

type NervousSystem interface {
	Version()
}

type Nervous struct{}

func (this *Nervous) Version() {
	fmt.Println("This is descibe NervousSystem basic on Tcp communicate")
}
