package main

import "fmt"

type Node struct {
	value string,
	next *Node
}

type MixedStack []Node

func (m *MixedStack) Push(value string) {
	*m = append(*m, )
}

/*Describe how you could use a single array to implement three stacks*/
/*@TODO*/
func main() {
	stack := Stack{}
	stack.Push("item1")
	stack.Push("item2")
	stack.Push("item3")

	fmt.Printf("%+v\n", stack)
}
