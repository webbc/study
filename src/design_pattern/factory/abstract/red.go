package abstract

import "fmt"

type Red struct {
}

func (r *Red) SayColor() {
	fmt.Println("red")
}
