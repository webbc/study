package main

import "design_pattern/factory/abstract"

func main() {
	ff := &abstract.FruitFactory{}
	apple := ff.GetFruit(abstract.FRUIT_APPLE)
	apple.SayPrice()

	banana := ff.GetFruit(abstract.FRUIT_BANANA)
	banana.SayPrice()

	cf := &abstract.ColorFactory{}
	red := cf.GetColor(abstract.COLOR_RED)
	red.SayColor()

	yellow := cf.GetColor(abstract.COLOR_YELLOW)
	yellow.SayColor()
}
