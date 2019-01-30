package template_function

import "fmt"

type EatRice struct {
	son IEatRice
}

func (e *EatRice) Eat() {
	e.start()
	e.son.eat()
	e.end()
}

func (e *EatRice) start() {
	fmt.Println("点餐")
}

func (e *EatRice) end() {
	fmt.Println("付账")
}

func NewEatRice(son IEatRice) *EatRice {
	return &EatRice{
		son: son,
	}
}
