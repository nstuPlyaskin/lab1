package main

import (
	"fmt"

	"lab1/src/structs"
)

func main() {
	stack := &structs.Stack{}
	stack.Spush("1asd")
	stack.Spush("2asd")
	stack.Spush("3asd")

	queue := &structs.Queue{}
	queue.Qpush("1aabo")
	queue.Qpush("2aabo")
	queue.Qpush("3aabo")

	err, val := stack.Spop()
	fmt.Println(err, val)
	err, val = stack.Spop()
	fmt.Println(err, val)
	err, val = stack.Spop()
	fmt.Println(err, val)
	err, val = stack.Spop()
	fmt.Println(err, val)
	err, val = stack.Spop()
	fmt.Println(err, val)
	err, val = stack.Spop()
	fmt.Println(err, val)
	err, val = stack.Spop()
	fmt.Println(err, val)

	err, val = queue.Qpop()
	fmt.Println(err, val)

	curentNode := stack.Head
	for curentNode != nil {
		fmt.Println(curentNode.Data)
		curentNode = curentNode.Next
	}

	curentNode1 := queue.Head
	for curentNode1 != nil {
		fmt.Println(curentNode1.Data)
		curentNode1 = curentNode1.Next
	}
}
