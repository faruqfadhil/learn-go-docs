package main

import (
	"fmt"
	"strings"
)

type SingleLinkedList struct {
	Head *Node
}
type Node struct {
	Data int
	Next *Node
}

func (s *SingleLinkedList) Print() string {
	temp := s.Head
	out := []string{}
	for temp != nil {
		out = append(out, fmt.Sprintf("%d", temp.Data))
		temp = temp.Next
	}
	return strings.Join(out, "->")
}

func (s *SingleLinkedList) InsertFront(data int) {
	newNode := &Node{
		Data: data,
	}
	newNode.Next = s.Head
	s.Head = newNode
}

func (s *SingleLinkedList) InsertAfter(node *Node, data int) {
	newNode := &Node{
		Data: data,
	}
	newNode.Next = node.Next
	node.Next = newNode
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

func (s *SingleLinkedList) Delete(node *Node) {
	if s.Head.Data == node.Data {
		s.Head = node.Next
		return
	}

	prevNode := s.Head
	for prevNode.Next.Data != node.Data {
		prevNode = prevNode.Next
	}
	prevNode.Next = node.Next
}
