package main

import (
	"design_pattern/strategy"
)

func main() {

	w1 := &strategy.Work{
		Tool: &strategy.CarStrategy{},
	}

	w1.GoToWork()

	w2 := &strategy.Work{
		Tool: &strategy.BusStrategy{},
	}

	w2.GoToWork()

}
