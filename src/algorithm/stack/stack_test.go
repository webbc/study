package stack

import (
	"testing"
	"fmt"
)

func TestStack(t *testing.T) {
	stack := newStack(5)
	for i := 1; i < 10; i++ {
		stack.Push(i)
	}
	fmt.Println(stack)
	for i := 1; i < 10; i++ {
		data := stack.Pop()
		if data == nil {
			fmt.Println("data is nil")
		} else {
			fmt.Println(data)
		}
		fmt.Println(stack)
	}
	for i := 1; i < 10; i++ {
		stack.Push(i*2)
	}
	fmt.Println(stack)
}
