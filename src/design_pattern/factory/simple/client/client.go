package main

import "design_pattern/factory/simple"

func main() {
	simpleFactory := &simple.FruitFactory{}
	apple := simpleFactory.Get(simple.FRUIT_APPLE)
	apple.SayPrice()

	banana := simpleFactory.Get(simple.FRUIT_BANANA)
	banana.SayPrice()
}
