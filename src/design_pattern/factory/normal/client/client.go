package main

import "design_pattern/factory/normal"

func main() {
	ff := &normal.FruitFactory{}
	apple := ff.GetFruit(normal.FRUIT_APPLE)
	apple.SayPrice()

	banana := ff.GetFruit(normal.FRUIT_BANANA)
	banana.SayPrice()
}
