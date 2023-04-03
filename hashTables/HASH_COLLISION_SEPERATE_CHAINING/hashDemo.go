package main

import (
	"fmt"
)

// size of the array.
const size = 10

// hashTable
type hashTable struct {
	array [size]*bucket
}

// linked list
type bucket struct {
	head *bucketNode
}

// each node
type bucketNode struct {
	key  string
	next *bucketNode
}

// insertion to linkedlist
func (h *hashTable) insertionHashTable(key string) {
	index := hash(key)
	h.array[index].insertionBucket(key)
}

// insertion
func (b *bucket) insertionBucket(key string) {
	if !b.searchingBucket(key) {
		newNode := new(bucketNode)
		newNode.key = key
		newNode.next = b.head
		b.head = newNode
		fmt.Println("Inserted")
	} else {
		fmt.Println("Already exists!")
	}

}

// hashing
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % size
}

// searching
func (h *hashTable) searching(key string) bool {
	index := hash(key)
	return h.array[index].searchingBucket(key)
}

// searching bucket
func (b *bucket) searchingBucket(key string) bool {
	temp := b.head
	for temp != nil {
		if temp.key == key {
			return true
		}
		temp = temp.next
	}
	return false
}

// deleting
func (h *hashTable) delete(key string) {
	index := hash(key)
	h.array[index].deleteNode(key)
}

// deleting Node
func (b *bucket) deleteNode(key string) {
	if b.head.key == key {
		b.head = b.head.next
		fmt.Println("First element deleted")
		return
	}
	temp := b.head
	for temp.next != nil {
		if temp.next.key == key {
			temp.next = temp.next.next
			fmt.Println("deleted")
			return
		}
		temp = temp.next
	}
}

// main
func main() {
	fmt.Println("Seperate chaining example")
	myhashTable := hashTable{}
	fmt.Println(myhashTable)
	for i := range myhashTable.array {
		myhashTable.array[i] = &bucket{}
	}

	// list := [5]string{
	// 	"hello",
	// 	"ganesh",
	// 	"stebin",
	// 	"ortho",
	// 	"beta",
	// }
	// for _, v := range list {
	// 	myhashTable.insertionHashTable(v)
	// }

	myhashTable.insertionHashTable("cd")
	myhashTable.insertionHashTable("be")
	fmt.Println(myhashTable.searching("hello"))
	myhashTable.delete("be")
	fmt.Println(myhashTable.searching("cd"))
}
