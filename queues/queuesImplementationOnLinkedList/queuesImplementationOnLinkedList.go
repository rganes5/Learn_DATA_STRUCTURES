package main

import "fmt"

type Queues struct {
	head *node
	tail *node
}

type node struct {
	data int
	next *node
}

func main() {
	myQueues := Queues{}
	fmt.Println("Queue implementation on linkedlist")
	var choice int
loop:
	for {
		fmt.Print("\nEnter\n1 for ENQUEUE\n2 for DEQUEUE\n3 for DISPLAY\n0 to EXIT\n")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			myQueues.enqueue()
		case 2:
			myQueues.dequeue()
		case 3:
			myQueues.display()
		case 0:
			break loop
		default:
			fmt.Println("Enter the correct option")
		}
	}
}

func (q *Queues) enqueue() {
	var data int
	fmt.Println("Enter the values")
	fmt.Scan(&data)
	newNode := new(node)
	newNode.data = data
	newNode.next = nil
	if q.head == nil {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}

}

func (q *Queues) dequeue() {
	if q.head == nil {
		fmt.Println("QUEUE UNDERFLOW")
		return
	} else {
		q.head = q.head.next

	}
}

func (q *Queues) display() {
	temp := q.head
	if q.head == nil {
		fmt.Println("Queue empty")
	} else {
		for temp != nil {
			fmt.Println(temp.data)
			temp = temp.next
		}
	}
}
