package main

import (
	"design_pattern/cor"
	"math/rand"
)

func main() {
	h := cor.CreateHandler()

	for i := 1; i <= 100; i++ {
		price := rand.Float64()
		h.PriceHandle(price)
	}

}
