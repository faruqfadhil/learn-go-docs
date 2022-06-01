package main

import (
	"fmt"
	"strings"
)

type Queue struct {
	Queue *QueueNode
}

type QueueNode struct {
	Data int
	Next *QueueNode
}

func (q *Queue) Enqueue(data int) {
	newData := &QueueNode{
		Data: data,
	}
	if q.Queue == nil {
		q.Queue = newData
		return
	}

	temp := q.Queue
	for temp.Next != nil {
		temp = temp.Next
	}
	temp.Next = newData
}

func (q *Queue) Dequeue() {
	if q.Queue == nil {
		fmt.Println("empty queue")
	}

	q.Queue = q.Queue.Next
}

func (q *Queue) Print() string {
	temp := q.Queue
	out := []string{}
	for temp != nil {
		out = append(out, fmt.Sprintf("%d", temp.Data))
		temp = temp.Next
	}
	return strings.Join(out, ",")
}

func main() {
	queue := &Queue{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Dequeue()
	queue.Dequeue()
	queue.Enqueue(30)
	queue.Enqueue(40)
	queue.Enqueue(50)
	queue.Dequeue()
	fmt.Println(queue.Print())
}
