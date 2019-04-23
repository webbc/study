package simple

import "fmt"

// 苹果
type Apple struct {
}

func (a *Apple) SayPrice() {
	fmt.Printf("苹果的单价:1元\n")
}
