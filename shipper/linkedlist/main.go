package main

import (
	"fmt"
)

type SingleLinkedList struct {
	Head *Node
}
type Node struct {
	Data int
	Next *Node
}

func (s *SingleLinkedList) Append(data int) {
	newNode := &Node{
		Data: data,
	}
	newNode.Next = nil

	if s.Head == nil {
		s.Head = newNode
		return
	}

	temp := s.Head
	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = newNode
}

func (s *SingleLinkedList) getCount() int {
	var count int
	current := s.Head
	for current != nil {
		fmt.Println(current.Data)
		count++
		current = current.Next
	}
	return count
}

func main() {
	var testCases int
	fmt.Scan(&testCases)
	for i := 0; i < testCases; i++ {
		init := &SingleLinkedList{Head: &Node{}}
		var N int
		fmt.Scan(&N)

		arr := make([]int, N)
		for i := 0; i < N; i++ {
			if _, err := fmt.Scan(&arr[i]); err != nil {
				panic(err)
			}
			init.Append(arr[i])
		}
		fmt.Println(init.getCount() - 1)
	}
}
