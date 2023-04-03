//STACK IMPLEMENTATION ON ARRAYS

package main

import "fmt"

type stack struct {
	items []int
}

func main() {
	var choice int
	myStack := stack{}
	fmt.Println("Empty stack")
	fmt.Println(myStack)
	fmt.Println("Stack implementation on Arrays")
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
	var items int
	fmt.Println("Enter the values: ")
	fmt.Scan(&items)
	s.items = append(s.items, items)
}

func (s *stack) display() {
	fmt.Println(s.items)
}

func (s *stack) pop() {
	if len(s.items) == 0 {
		fmt.Println("Stack Underflow")
		return
	}
	temp := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	fmt.Printf("%v is removed Successfully\n", temp)
}
