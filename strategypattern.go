package main

import "fmt"

type Order string

const (
	kill         = "kill him/she"
	ordertomercy = "mercy him/she"
)

type Words interface {
	sayWord() Order
}

type King struct {
	say Order
}

func (k *King) sayWord() Order {
	return k.say
}

type SayOrder struct {
}

func (s *SayOrder) sayorder(w Words) {
	fmt.Println(fmt.Sprintf("I said %s !!!!", w.sayWord()))
}

func main() {
	so := &SayOrder{}
	k := &King{kill}
	k1 := &King{ordertomercy}
	so.sayorder(k)
	so.sayorder(k1)
}
