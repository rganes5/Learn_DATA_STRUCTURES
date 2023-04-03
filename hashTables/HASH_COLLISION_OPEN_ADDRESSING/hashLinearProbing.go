package main

import "fmt"

const size = 10

type hashTable struct {
	array [size]string
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % size
}

// LINEAR PROBING

func (h *hashTable) linearProbing(key string) {
	index := hash(key)
	fmt.Println(index)
	if h.array[index] == " " {
		h.array[index] = key
	} else {
		for i := 1; i < size; i++ {
			if h.array[index+i] == " " {
				h.array[index+i] = key
				return
			}
		}
	}
}

//QUADRACTIC PROBING
func (h *hashTable) quadracticProbing(key string) {
	index := hash(key)
	if h.array[index] == " " {
		h.array[index] = key
	} else {
		for i := 1; i < size; i++ {
			if h.array[index+(i*i)] == " " {
				h.array[index+(i*i)] = key
				return
			}
		}
	}
}

//DOUBLE HASHING

func main() {
	myHashTable := hashTable{}
	fmt.Println("Before")
	fmt.Println(myHashTable)
	fmt.Println("After")
	//
	//LINEAR PROBING COLLISION DEMONSTRATION
	myHashTable.linearProbing("hello")
	myHashTable.linearProbing("hello")
	fmt.Println(myHashTable)
	//
	//QUADRATIC EQUATION COLLISION DEMONSTRATION
	myHashTable.quadracticProbing("hello")
	fmt.Println(myHashTable)

}
