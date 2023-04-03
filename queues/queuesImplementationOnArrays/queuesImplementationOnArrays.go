package main

import "fmt"

type Queues struct {
	data []int
}

func main() {
	var choice int
	fmt.Println("Queues implementation on Arrays")
	myQueues := Queues{}
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

//ENQUEUE
func (q *Queues) enqueue() {
	var data int
	fmt.Println("Enter the values: ")
	fmt.Scan(&data)
	q.data = append(q.data, data)
}

//DEQUEUE
func (q *Queues) dequeue() {
	if len(q.data) == 0 {
		fmt.Println("QUEUE Underflow")
	} else {
		temp := q.data[0]
		q.data = q.data[1:]
		fmt.Printf("\n%v is deleted\n", temp)
	}
}

func (q *Queues) display() {
	if len(q.data) == 0 {
		fmt.Println("QUEUE EMPTY!!")
	}
	fmt.Println(q.data)
}
