package stack

import "errors"

type Stack struct {
	cap  int           // 容量
	top  int           // 栈顶
	data []interface{} // 数据
}

func newStack(cap int) *Stack {
	return &Stack{
		cap:  cap,
		top:  0,
		data: make([]interface{}, cap, cap),
	}
}

// 判断栈是否已满
func (s *Stack) isFull() bool {
	if s.top == s.cap {
		return true
	}
	return false
}

// 判断栈是否为空
func (s *Stack) isEmpty() bool {
	if s.top == 0 {
		return true
	}
	return false
}

// 出栈
func (s *Stack) Pop() interface{} {
	if !s.isEmpty() {
		value := s.data[s.top-1]
		s.data[s.top-1] = nil
		s.top--
		return value
	}
	return nil
}

// 入栈
func (s *Stack) Push(value interface{}) error {
	if s.isFull() {
		return errors.New("stack is full")
	}
	s.data[s.top] = value
	s.top++
	return nil
}
