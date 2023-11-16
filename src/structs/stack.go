package structs

import "errors"

type Stack struct {
	Head *Node
}

type Node struct {
	Data string
	Next *Node
}

func (stack *Stack) SPUSH(value string) error {
	if value == "" {
		return errors.New("Error: Stack is empty!")
	}
	newNode := &Node{Data: value}
	if stack.Head == nil {
		stack.Head = newNode
	} else {
		newNode.Next = stack.Head
		stack.Head = newNode
	}
	return nil
}

func (stack *Stack) SPOP() (error, string) {
	if stack.Head == nil {
		return errors.New("Error: can't find head."), ""
	}

	temp := stack.Head.Data
	stack.Head = stack.Head.Next

	return nil, temp
}
