package main

import (
	"fmt"
)

//SPECIFYING NODE AND LIST

type node struct {
	data int
	next *node
}

type singlelinkedlist struct {
	head *node
	tail *node
}

//MAIN FUNCTION

func main() {
	var mainChoice int
	fmt.Println("Welcome to Linked List Basic Operations :) ")
	list := singlelinkedlist{}

loop:
	for {
		fmt.Println("Enter 1 Array to LinkedList Conversions")
		fmt.Println("Enter 2 for basic Linked List Operations")
		fmt.Println("Enter 0 for EXIT")
		fmt.Scan(&mainChoice)
		switch mainChoice {
		case 1:
			list.convert()
		case 2:
			list.operations()
		case 0:
			break loop
		default:
			fmt.Println("Enter correct option.")

		}
	}
}

func (list *singlelinkedlist) operations() {
	var choice int
loop:
	for {
		fmt.Println("Enter which operation you would like to perform: ")
		fmt.Printf("\n1--> Insert a element from the Last\n2--> Insert a element in the beginning\n3--> Display\n4--> Delete\n5--> Delete specific element\n6--> Sort\n7--> Delete duplicates\n8--> Display Reverse \n0--> TO  EXIT\n")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			list.insert(0)
		case 2:
			list.insertBeg()
		case 3:
			list.display()
		case 4:
			list.delete()
		case 5:
			list.deletespecific()
		case 6:
			list.sort()
		case 7:
			list.removeDuplicates()
		case 8:
			list.displayReverse()
		case 0:
			break loop
		default:
			fmt.Println("Enter the correct option")

		}
	}

}

//INSERTION

func (list *singlelinkedlist) insert(number int) {
	var data int
	if number == 0 {
		fmt.Println("Enter the number: ")
		fmt.Scan(&data)
	} else {
		data = number
	}
	newNode := new(node)
	newNode.data = data
	newNode.next = nil

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		list.tail.next = newNode
		list.tail = newNode
	}
}

// INSERT BEGINNING

func (list *singlelinkedlist) insertBeg() {
	var data int
	fmt.Println("Enter the number: ")
	fmt.Scan(&data)
	newNode := new(node)
	newNode.data = data
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.next = list.head
		list.head = newNode
	}
}

//DISPLAY

func (list *singlelinkedlist) display() {
	if list.head == nil {
		fmt.Println("The list is empty")
	} else {
		fmt.Println("The list is:  ")
		temp := list.head
		for temp != nil {
			fmt.Printf("%v\t", temp.data)
			temp = temp.next
		}
		fmt.Printf("\n")
	}
}

// DISPLAY REVERSE

func (list *singlelinkedlist) displayReverse() {
	if list.head == nil {
		fmt.Println("The list is empty")
	} else {
		fmt.Println("The list is:  ")
		temp := list.head
		for temp != nil {
			fmt.Printf(" %v ", temp.data)
			temp = temp.next
		}
		fmt.Printf("\n")
	}
}

// DELETE
func (list *singlelinkedlist) delete() {
	if list.head == nil {
		fmt.Println("The list is empty")
	} else {
		temp := list.head
		for temp != nil {
			if temp.next.next == nil {
				temp.next = nil
			}
			temp = temp.next
		}
	}
}

// DELETE SPECIFIC
func (list *singlelinkedlist) deletespecific() {
	var specific int
	fmt.Println("Enter the element that you would like to delete: ")
	fmt.Scan(&specific)
	if list.head == nil {
		fmt.Println("The list is empty")
	} else {
		temp := list.head
		if temp.data == specific {
			list.head = temp.next
		}
		for temp != nil {
			if temp.next.data == specific {
				temp.next = temp.next.next
				break
			}
			temp = temp.next
		}
	}
}

// SORT
func (list *singlelinkedlist) sort() {
	if list.head == nil {
		fmt.Println("The list is empty")
	} else {
		temp1 := list.head
		for temp1 != nil {
			temp2 := list.head
			for temp2 != nil {
				if temp1.data < temp2.data {
					t := temp1.data
					temp1.data = temp2.data
					temp2.data = t
				}
				temp2 = temp2.next
			}
			temp1 = temp1.next
		}
	}
}

// REMOVE DUPLICATES
func (list *singlelinkedlist) removeDuplicates() {
	if list.head == nil {
		fmt.Println("The list is empty")
	} else {

		key := list.head
		for key != nil {
			prev := key
			temp := key.next
			for temp != nil {
				if key.data == temp.data {
					prev.next = temp.next
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

// CONVERT ARRAY TO LINKED LIST
func (list *singlelinkedlist) convert() {
	var userSlice = []int{}
	var limit, val int
	fmt.Println("Enter the array limit")
	fmt.Scan(&limit)
	fmt.Println("Enter the array elements")
	for i := 0; i < limit; i++ {
		fmt.Scan(&val)
		userSlice = append(userSlice, val)
	}
	fmt.Printf("\nThe entered array is :\t %v ", userSlice)
	//conversion
	for _, number := range userSlice {
		list.insert(number)
	}
	fmt.Println("After conversion :")
	list.display()

}
