package abstract

import "fmt"

type Yellow struct {
}

func (y *Yellow) SayColor() {
	fmt.Println("yellow")
}
