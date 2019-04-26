package main

import (
	"design_pattern/memento"
	"fmt"
)

func main() {
	c := &memento.CareTaker{}

	o := &memento.Originator{}
	o.SetState(1)
	m := o.Save()
	c.Append(m)

	o.SetState(2)
	m = o.Save()
	c.Append(m)

	o.SetState(3)
	fmt.Printf("current state:%v\n", o.GetState())
	fmt.Printf("state:%v\n", c.Get(0).GetState())
	fmt.Printf("state:%v\n", c.Get(1).GetState())
}
