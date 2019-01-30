package strategy

import "fmt"

type BusStrategy struct {
}

func (b *BusStrategy) Do() {
	fmt.Println("坐公交")
}
