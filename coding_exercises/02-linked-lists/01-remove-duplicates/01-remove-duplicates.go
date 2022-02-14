package main

import (
	"fmt"
	"os"
)

type Node struct {
	value string
	next  *Node
}

func (n *Node) ToString() string {
	nextVal := "nil"
	if n.next != nil {
		nextVal = n.next.value
	}
	return fmt.Sprintf("{value: %s, next: %s}", n.value, nextVal)
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) add(value string) Node {
	var newNode Node = Node{value, nil}
	if l.head == nil {
		l.head = &newNode
	}
	if l.tail != nil {
		l.tail.next = &newNode
	}
	l.tail = &newNode

	return newNode
}

func (l *LinkedList) getNode(search string) *Node {
	cursor := l.head
	for search != cursor.value {
		fmt.Printf("searchin if element %s is equal to %s\n", cursor.ToString(), search)
		if cursor.next == nil {
			fmt.Println("[error] element not found in linked list")
			os.Exit(1)
		}
		cursor = cursor.next
	}
	fmt.Printf("found element with value %s : %s\n", search, cursor.ToString())
	return cursor
}

func (l *LinkedList) remove(search string) {
	current := l.getNode(search)
	temp := l.head
	if current == nil {
		fmt.Printf("Error, no element found with value %s", search)
		return
	}
	if l.head == current {
		l.head = temp.next
		return
	}
	for temp.next != current {
		temp = temp.next
	}
	if l.tail == current {
		l.tail = temp
		l.tail.next = nil
		return
	}
	if temp.next == current {
		temp.next = current.next
	}
}

func (l *LinkedList) print(debug bool) {
	elem := l.head
	fmt.Println("all items:")
	for elem != nil {
		if debug {
			var flags string
			if elem == l.head {
				flags = flags + "H"
			}
			if elem == l.tail {
				flags = flags + "T"
			}
			fmt.Printf("[%s]%s\n", flags, elem.value)
		} else {
			fmt.Println(elem.value)
		}
		elem = elem.next
	}
}

func main() {
	debug := true
	myList := LinkedList{}
	myList.add("item1")
	myList.add("item2")
	myList.add("item3")
	myList.add("item4")
	myList.add("item5")
	myList.print(debug)

	fmt.Println("------ remove item 4 ------")
	myList.remove("item4")
	myList.print(debug)
	//como fazer uma lista encadeada em golang?
	fmt.Println("------ remove item 1 ------")
	myList.remove("item1")
	myList.print(debug)

	fmt.Println("------ remove item 5 ------")
	myList.remove("item5")
	myList.print(debug)
}
