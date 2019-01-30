package main

import "design_pattern/template_function"

func main() {
	// 创建A号顾客
	a := &template_function.CustomerA{}

	// a吃饭
	template_function.NewEatRice(a).Eat()

	// 创建B号顾客
	b := &template_function.CustomerB{}

	// b吃饭
	template_function.NewEatRice(b).Eat()
}
