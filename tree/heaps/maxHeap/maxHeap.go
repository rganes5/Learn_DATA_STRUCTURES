package main

import "fmt"

type maxHeap struct {
	array []int
}

// Inserting into heap
func (m *maxHeap) insert(value int) {
	m.array = append(m.array, value)
	m.heapifyBotToTop()
}

// Rearranging the elements from bottom to top to satisfy the heap property
func (m *maxHeap) heapifyBotToTop() {
	index := len(m.array) - 1
	for m.array[parent(index)] < m.array[index] {
		m.array[parent(index)], m.array[index] = m.array[index], m.array[parent(index)]
		index = parent(index)
	}
}

// deleting the root node
func (m *maxHeap) extractRoot() {
	element := m.array[0]
	m.array[0] = m.array[len(m.array)-1]
	m.array = m.array[:len(m.array)-1]
	fmt.Println(element)
	m.heapifyTopToBot()
}

//rearranging the elements from top to bottom to satisfy the heap property
func (m *maxHeap) heapifyTopToBot() {
	index := 0
	for lChild(index) <= len(m.array)-1 {
		lc := lChild(index)
		rc := rChild(index)
		if rc > len(m.array)-1 {
			if m.array[lc] > m.array[index] {
				m.array[index], m.array[lc] = m.array[lc], m.array[index]
				index = lChild(index)
			} else {
				return
			}
		} else if m.array[lc] > m.array[rc] && m.array[lc] > m.array[index] {
			m.array[index], m.array[lc] = m.array[lc], m.array[index]
			index = lChild(index)
		} else if m.array[rc] > m.array[lc] && m.array[rc] > m.array[index] {
			m.array[index], m.array[rc] = m.array[rc], m.array[index]
			index = rChild(index)
		} else {
			return
		}
	}
}

// To find the parent node of a node
func parent(index int) int {
	return (index - 1) / 2
}

// To find the left child of a node
func lChild(index int) int {
	return (index * 2) + 1
}

// To find the right child of a node
func rChild(index int) int {
	return (index * 2) + 2
}

func main() {
	var m = maxHeap{}
	sampleArray := []int{1, 5, 2, 6, 3, 4}
	for _, v := range sampleArray {
		m.insert(v)
	}
	fmt.Printf("\nMax heap:%v\n", m.array)
	m.extractRoot()
	fmt.Printf("\nMax heap:%v\n", m.array)

}
