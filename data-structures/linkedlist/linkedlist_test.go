package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrint(t *testing.T) {
	tests := map[string]struct {
		req *SingleLinkedList
		out string
		err error
	}{
		"success": {
			req: &SingleLinkedList{&Node{
				Data: 1,
				Next: &Node{
					Data: 2,
					Next: &Node{
						Data: 3,
					},
				},
			},
			},
			out: "1->2->3",
			err: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			out := test.req.Print()
			diff := cmp.Diff(test.out, out)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestInsertFront(t *testing.T) {
	tests := map[string]struct {
		init *SingleLinkedList
		req  int
		out  string
		err  error
	}{
		"success": {
			init: &SingleLinkedList{
				Head: &Node{
					Data: 1,
					Next: &Node{
						Data: 2,
						Next: &Node{
							Data: 3,
						},
					},
				},
			},
			req: 5,
			out: "5->1->2->3",
			err: nil,
		},
		"success with nil": {
			init: &SingleLinkedList{
				Head: nil,
			},
			req: 5,
			out: "5",
			err: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.init.InsertFront(test.req)
			out := test.init.Print()
			diff := cmp.Diff(test.out, out)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestInsertAfter(t *testing.T) {
	tests := map[string]struct {
		init *SingleLinkedList
		req  int
		out  string
		err  error
	}{
		"success": {
			init: &SingleLinkedList{
				Head: &Node{
					Data: 1,
					Next: &Node{
						Data: 2,
						Next: &Node{
							Data: 3,
						},
					},
				},
			},
			req: 5,
			out: "1->2->5->3",
			err: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.init.InsertAfter(test.init.Head.Next, test.req)
			out := test.init.Print()
			diff := cmp.Diff(test.out, out)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	tests := map[string]struct {
		init *SingleLinkedList
		req  int
		out  string
		err  error
	}{
		"success": {
			init: &SingleLinkedList{
				Head: &Node{
					Data: 1,
					Next: &Node{
						Data: 2,
						Next: &Node{
							Data: 3,
						},
					},
				},
			},
			req: 5,
			out: "1->2->3->5",
			err: nil,
		},
		"success with nil": {
			init: &SingleLinkedList{
				Head: nil,
			},
			req: 5,
			out: "5",
			err: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			test.init.Append(test.req)
			out := test.init.Print()
			diff := cmp.Diff(test.out, out)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type requestDelete string
	const (
		middle requestDelete = "middle"
		head   requestDelete = "head"
	)
	tests := map[string]struct {
		init *SingleLinkedList
		req  requestDelete
		out  string
		err  error
	}{
		"success": {
			init: &SingleLinkedList{
				Head: &Node{
					Data: 1,
					Next: &Node{
						Data: 2,
						Next: &Node{
							Data: 3,
						},
					},
				},
			},
			req: middle,
			out: "1->3",
			err: nil,
		},
		"success with head": {
			init: &SingleLinkedList{
				Head: &Node{
					Data: 1,
					Next: &Node{
						Data: 2,
						Next: &Node{
							Data: 3,
						},
					},
				},
			},
			req: head,
			out: "2->3",
			err: nil,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			switch test.req {
			case middle:
				test.init.Delete(test.init.Head.Next)
			case head:
				test.init.Delete(test.init.Head)
			}

			out := test.init.Print()
			diff := cmp.Diff(test.out, out)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
