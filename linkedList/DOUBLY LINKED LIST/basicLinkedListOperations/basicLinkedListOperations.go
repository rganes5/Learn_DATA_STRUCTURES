package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to doubly linked list ")
	list := doublyLinkedList{}
	list.start()
}

type node struct {
	data int
	next *node
	prev *node
}

type doublyLinkedList struct {
	head *node
	tail *node
}

func (list *doublyLinkedList) start() {
	var choice int
loop:

	for {
		fmt.Printf("Enter \n1 to convert array to a Doubly Linked List\n2 to perform LinkedList Operations\n0 to EXIT!\n")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			list.convert()
		case 2:
			list.operations()
		case 0:
			break loop
		default:
			fmt.Println("Wrong Option")
		}
	}
}

// BASIC OPERATIONS
func (list *doublyLinkedList) operations() {
	var choice int
loop:
	for {
		fmt.Println("1--> Insertion\n2--> Insert at Beginning\n3--> Display\n4--> Delete last node\n5--> Delete Specific\n6--> SORT\n7--> Delete Duplicate\n8--> Reverse\n0--> EXIT")
		fmt.Println("Choose a option")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			list.insert(0)
		case 2:
			list.insertBeg()
		case 3:
			list.display()
		case 4:
			list.deleteLast()
		case 5:
			list.deleteSpecific()
		case 6:
			list.sort()
		case 7:
			list.deleteDuplicate()
		case 8:
			list.reverse()
		case 0:
			break loop
		default:
			fmt.Println("Wrong option")
		}
	}
}

// INSERT AT END
func (list *doublyLinkedList) insert(v int) {
	var data int
	if v == 0 {
		fmt.Println("Enter the values")
		fmt.Scan(&data)
	} else {
		data = v
	}

	newNode := new(node)
	newNode.data = data
	newNode.next = nil
	newNode.prev = nil
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		newNode.prev = list.tail
		list.tail = newNode
	}
}

// INSERT AT BEGINNING
func (list *doublyLinkedList) insertBeg() {
	var data int
	fmt.Println("Enter the values")
	fmt.Scan(&data)
	newNode := new(node)
	newNode.data = data
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head.prev = newNode
		list.head = newNode
	}
}

// DISPLAY
func (list *doublyLinkedList) display() {
	if list.head == nil {
		fmt.Println("The list is empty...........!")
	} else {
		fmt.Println("The Linked List is")
		temp := list.head
		for temp != nil {
			fmt.Printf(" %v ", temp.data)
			temp = temp.next
		}
		fmt.Printf("\n")
	}
}

// DELETE LAST
func (list *doublyLinkedList) deleteLast() {
	if list.head == nil {
		fmt.Println("List is empty")
	} else {
		if list.head == list.tail {
			list.head = nil
			list.tail = nil
		} else {
			list.tail = list.tail.prev
			list.tail.next = nil
		}
		fmt.Println("deleted")
	}
}

// DELETE SPECIFIC
func (list *doublyLinkedList) deleteSpecific() {
	var specific int
	fmt.Println("Enter the element which you would like to delete :")
	fmt.Scan(&specific)
	if list.head == nil {
		fmt.Println("List is empty")
	} else {

		temp := list.head
		if temp.data == specific {
			list.head = temp.next
			list.head.prev = nil
		} else {
			for temp != nil {
				if temp.next.data == specific {
					temp.next = temp.next.next
					temp.next.prev = temp
					break
				}
				temp = temp.next

			}
		}
	}
	fmt.Println("deleted")
}

// SORT
func (list *doublyLinkedList) sort() {
	if list.head == nil {
		fmt.Println("Nothing to Sort")
	} else {
		key := list.head
		for key != nil {
			temp := list.head
			for temp != nil {
				if key.data < temp.data {
					t := key.data
					key.data = temp.data
					temp.data = t
				}
				temp = temp.next
			}
			key = key.next
		}
	}
	fmt.Println("Sorted")
}

// DELETE DUPLICATE
func (list *doublyLinkedList) deleteDuplicate() {
	if list.head == nil {
		fmt.Println("There is no data to delete...........!")
	} else {
		key := list.head
		for key != nil {
			prev := key
			temp := key.next
			for temp != nil {
				if temp.data == key.data {
					prev.next = temp.next
					temp.next.prev = prev
					if temp == list.tail {
						list.tail = prev
					}
				}
				prev = temp
				temp = temp.next
			}
			key = key.next
		}

	}
}

// REVERSE
func (list *doublyLinkedList) reverse() {
	if list.head == nil {
		fmt.Println("List is empty")
	} else {
		fmt.Println("The list reversed is")
		temp := list.tail
		for temp != nil {
			fmt.Printf(" %v ", temp.data)
			temp = temp.prev
		}
		println()
	}
}

// CONVERT ARRAY TO DOUBLY LINKED LIST
func (list *doublyLinkedList) convert() {
	var userSlices = []int{}
	var limit, values int
	fmt.Println("Enter the array limit")
	fmt.Scan(&limit)
	fmt.Println("Enter the array elements")
	for i := 0; i < limit; i++ {
		fmt.Scan(&values)
		userSlices = append(userSlices, values)
	}
	fmt.Printf("The array is :%v\n", userSlices)
	//CONVERSION
	for _, v := range userSlices {
		list.insert(v)
	}
	list.display()

}
