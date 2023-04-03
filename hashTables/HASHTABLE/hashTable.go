package main

import (
	"fmt"
)

// Array size is the size of the hashtable
const size = 5

// HASHTABLE to hold an array
type HashTable struct {
	array [size]*bucket
}

// bucket is a linked list in each slot of collision.
type bucket struct {
	head *bucketNode
}

// each node
type bucketNode struct {
	key  string
	next *bucketNode
}

// insert will take key and add it to array.
func (h *HashTable) insert(key string) {
	index := hash(key)
	h.array[index].insertValueToNode(key)
}

// hash converts the key and returns specific hash index
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % size
}

// insertvaluetonode will get the key, creates a node with the key and adds the node to the bucket.
func (b *bucket) insertValueToNode(key string) {
	if !b.searchBucket(key) {
		newbucketNode := new(bucketNode)
		newbucketNode.key = key
		newbucketNode.next = b.head
		b.head = newbucketNode
		fmt.Println(newbucketNode.key)
	} else {
		fmt.Println(key, "already exists")
	}

}

// SEARCH IN HASHTABLE
func (h *HashTable) searchHash(key string) bool {
	index := hash(key)
	return h.array[index].searchBucket(key)
}

// SEARCH IN BUCKET
func (b *bucket) searchBucket(key string) bool {
	temp := b.head
	for temp != nil {
		if key == temp.key {
			return true
		}
		temp = temp.next
	}
	return false
}

// DELETE FROM HASHTABLE
func (h *HashTable) deleteHash(key string) {
	index := hash(key)
	h.array[index].deleteBucket(key)
}

// DELETE FROM BUCKET
func (b *bucket) deleteBucket(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	temp := b.head
	for temp != nil {
		if temp.next.key == key {
			temp.next = temp.next.next
			fmt.Println("deleted")
		}
		temp = temp.next
	}
}

// MAIN FN
func main() {
	myHashTable := HashTable{}
	for i := range myHashTable.array {
		myHashTable.array[i] = &bucket{}
	}
	fmt.Println(myHashTable)
	list := []string{
		"Ganesh",
		"Stebin",
		"Edwin",
		"kyle",
		"Ramos",
	}
	for _, v := range list {
		myHashTable.insert(v)
	}
	myHashTable.deleteHash("kyle")
	fmt.Println(myHashTable)
	// myHashTable.insert("hello")
	// fmt.Println(myHashTable.searchHash("hello"))
	// myBucket := bucket{}
	// myBucket.insertValueToNode("hello")
	// fmt.Println(myBucket.search("hello"))
	// fmt.Println(myBucket.search("andi"))
}
