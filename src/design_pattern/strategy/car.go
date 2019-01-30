package strategy

import "fmt"

type CarStrategy struct {
}

func (c *CarStrategy) Do() {
	fmt.Println("开车")
}
