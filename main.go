package main

import (
	"fmt"
	"lab1/src/structs"
)

func main() {
	// STACK - PUSH
	stack := &structs.Stack{}
	stack.SPUSH("1asd")

	// STACK - POP
	err, val := stack.SPOP()
	fmt.Println(err, val)

	// QUEUE - PUSH
	queue := &structs.Queue{}
	queue.QPUSH("1aabo")

	// QUEUE - POP
	err, val = queue.QPOP()
	fmt.Println(err, val)

	// SET - ADD
	set := &structs.Set{}
	set.SADD("1")
	set.SADD("3")
	set.SADD("2")
	
	// SET - SREM
	err, val = set.SREM("body")
	fmt.Println(err, val)


	// OUTPUT INFO
	stackOutput := stack.Head
	for stackOutput != nil {
		fmt.Println(stackOutput.Data)
		stackOutput = stackOutput.Next
	}

	queueOutput := queue.Head
	for queueOutput != nil {
		fmt.Println(queueOutput.Data)
		queueOutput = queueOutput.Next
	}

	setOutput := set.Head
	for setOutput != nil {
		fmt.Println(setOutput.Data)
		setOutput = setOutput.Next
	}
}
