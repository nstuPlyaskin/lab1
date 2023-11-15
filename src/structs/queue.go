package structs

import "errors"

type Queue struct {
	Head *Node
	Tail *Node
}

func (queue *Queue) Qpush(value string) error {
	if value == "" {
		return errors.New("Error: Queue is empty.")
	}

	node := &Node{Data: value}
	if queue.Head == nil {
		queue.Head = node
		queue.Tail = node
	} else {
		queue.Tail.Next = node
		queue.Tail = node
	}
	return nil
}

func (queue *Queue) Qpop() (error, string) {
	if queue.Head == nil {
		return errors.New("Error: can't find head."), ""
	} else {
		tmp := queue.Head.Data
		queue.Head = queue.Head.Next
		return nil, tmp
	}
}