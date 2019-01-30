package queue

import (
	"errors"
	"fmt"
)

type Queue struct {
	head int           // 头角标
	tail int           // 尾角标
	cap  int           //容量
	data []interface{} // 数据
}

func NewQueue(cap int) *Queue {
	return &Queue{
		cap:  cap,
		head: 0,
		tail: 0,
		data: make([]interface{}, cap, cap),
	}
}

func (q *Queue) isFull() bool {
	if q.tail != 0 && q.tail-q.head == 0 {
		return true
	}
	return false
}

func (q *Queue) isEmpty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}

func (q *Queue) EnQueue(value interface{}) error {
	if q.isFull() {
		return errors.New("Queue is full")
	}
	if q.tail == q.cap {
		q.tail = 0
	}
	q.data[q.tail] = value
	q.tail++
	return nil
}

func (q *Queue) DeQueue() interface{} {
	if q.isEmpty() {
		fmt.Println("empty")
		return nil
	}
	if q.head == q.cap {
		q.head = 0
	}
	value := q.data[q.head]
	q.data[q.head] = nil
	q.head++
	return value
}
