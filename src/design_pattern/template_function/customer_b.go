package template_function

import "fmt"

type CustomerB struct {
	EatRice
}

func (c *CustomerB) eat() {
	fmt.Println("吃过桥米线")
}
