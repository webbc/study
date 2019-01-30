package template_function

import "fmt"

type CustomerA struct {
	EatRice
}

func (c *CustomerA) eat() {
	fmt.Println("吃麻辣香锅")
}
