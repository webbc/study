package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue(5)

	queue.DeQueue()

	queue.EnQueue(1)
	fmt.Println(queue)

	queue.EnQueue(2)
	fmt.Println(queue)

	queue.DeQueue()
	fmt.Println(queue)

	queue.DeQueue()
	fmt.Println(queue)

	queue.EnQueue(3)
	fmt.Println(queue)

	queue.EnQueue(4)
	fmt.Println(queue)

	queue.DeQueue()
	fmt.Println(queue)

	queue.EnQueue(5)
	fmt.Println(queue)

	queue.EnQueue(6)
	fmt.Println(queue)

	queue.EnQueue(7)
	fmt.Println(queue)
}
