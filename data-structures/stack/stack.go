package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	Stack *StackNode
}

type StackNode struct {
	Data int
	Next *StackNode
}

func (s *Stack) Push(data int) {
	temp := &StackNode{
		Data: data,
	}
	temp.Next = s.Stack
	s.Stack = temp
}

func (s *Stack) Pop() {
	if s.Stack != nil {
		s.Stack = s.Stack.Next
	}
}

func (s *Stack) Print() string {
	temp := s.Stack
	out := []string{}
	for temp != nil {
		out = append(out, fmt.Sprintf("%d", temp.Data))
		temp = temp.Next
	}
	return strings.Join(out, ",")
}

func main() {
	stack := &Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Pop()
	fmt.Println(stack.Print())
}
