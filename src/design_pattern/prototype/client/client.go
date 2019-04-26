package main

import (
	"design_pattern/prototype"
	"fmt"
)

func main() {
	bmw := &prototype.BMWCar{}
	bmw.SetName("bmw")
	bc := &prototype.BenChiCar{}
	bc.SetName("bc")

	clone_bmw := bmw.Clone()
	clone_bmw.SetName("clone_bmw")

	clone_bc := bc.Clone()
	clone_bc.SetName("clone_bc")

	fmt.Printf("bmw name:%v\n", bmw.GetName())
	fmt.Printf("bc name:%v\n", bc.GetName())
	fmt.Printf("clone_bmw name:%v\n", clone_bmw.GetName())
	fmt.Printf("clone_bc name:%v\n", clone_bc.GetName())
}
