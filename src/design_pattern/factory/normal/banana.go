package normal

import "fmt"

// 香蕉
type Banana struct {
}

func (b *Banana) SayPrice() {
	fmt.Printf("香蕉的单价:2元\n")
}
