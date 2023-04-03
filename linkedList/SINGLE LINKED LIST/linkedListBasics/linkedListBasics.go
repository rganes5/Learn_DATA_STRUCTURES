package main

import "fmt"

// Node represents a node of linked list
type node struct {
	data int
	next *node
}

// LinkedList represents a linked list
type linkedList struct {
	head   *node
	length int
}

//Inserting a new node in the first position

func (l *linkedList) prepend(value int) {
	newNode := node{data: value}
	// var n node
	// n.data=value
	if l.head != nil {
		newNode.next = l.head
		l.head = &newNode
		l.length++
	} else {
		l.head = &newNode
		l.length++
	}
	return
}

//Traversing and printing the linked list that we have
func (l *linkedList) printList() {
	if l.head == nil {
		return
	}

	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}

//Deleting a node
func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--

}
func main() {
	myList := linkedList{}
	myList.prepend(1)
	myList.prepend(2)
	myList.prepend(3)
	myList.prepend(4)
	myList.prepend(5)
	myList.prepend(6)
	myList.deleteWithValue(4)
	myList.printList()

}
