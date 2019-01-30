package slient

import "fmt"

type Car struct {
}

func (c *Car) drive() {
	fmt.Println("行驶中...")
}
