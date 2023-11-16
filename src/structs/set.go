package structs

import (
	"errors"
)

type Set struct {
	Head *Node
	Tail *Node
}

func(set *Set) SADD(value string) error{
	if value == "" {
		return errors.New("Error: Set is empty.")
	}

	newNode := &Node{Data: value}
	if set.Head == nil {
		set.Head = newNode
		set.Tail = newNode
	} else {
		set.Tail.Next = newNode
		set.Tail = newNode
	}

	return nil
}

func (set *Set) SREM(value string) (error, string) {
	if set.Head == nil {
		return errors.New("Error: set is empty."), ""
	}

	// обработка если первый элемент это искомый элемент
	if set.Head.Data == value {
		removedValue := set.Head.Data
		set.Head = set.Head.Next
		if set.Head == nil {
			set.Tail = nil
		}
		return nil, removedValue
	}

	// поиск
	current := set.Head
	for current.Next != nil && current.Next.Data != value {
		current = current.Next
	}

	// удаление
	if current.Next != nil {
		removedValue := current.Next.Data
		current.Next = current.Next.Next
		if current.Next == nil {
			set.Tail = current
		}
		return nil, removedValue
	}

	return errors.New("Error: element not found."), ""
}