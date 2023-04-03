package main

import (
	"fmt"
)

type stack struct {
	top *node //top is also head
}
type node struct {
	data int
	prev *node
}

func main() {
	var choice int
	fmt.Println("Stack implementation on Singly Linked List")
	myStack := stack{}
	// myStack.push(5)
	// fmt.Println(myStack)

loop:
	for {
		fmt.Print("Enter\n1 for PUSH\n2 for Display\n3 for POP\n0 to EXIT\n")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			myStack.push()
		case 2:
			myStack.display()
		case 3:
			myStack.pop()
		case 0:
			break loop
		default:
			fmt.Println("Invalid entry")
		}
	}

}

func (s *stack) push() {
	var data int
	fmt.Println("Enter the element to be inserted: ")
	fmt.Scan(&data)
	newnode := new(node)
	newnode.data = data
	newnode.prev = nil

	if s.top == nil {
		s.top = newnode
	} else {
		newnode.prev = s.top
		s.top = newnode
	}
}

func (s *stack) display() {
	if s.top == nil {
		fmt.Println("Stack is Empty")
	} else {
		fmt.Println("The stack is: ")
		temp := s.top
		for temp != nil {
			fmt.Println(temp.data)
			temp = temp.prev
		}
	}
}

func (s *stack) pop() {
	if s.top == nil {
		fmt.Println("Stack Underflow")
		return
	} else {
		temp := s.top.data
		s.top = s.top.prev
		fmt.Printf("%v is deleted", temp)
	}
}
