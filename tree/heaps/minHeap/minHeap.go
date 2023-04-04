package main

import "fmt"

type minHeap struct {
	array []int
}

//Inserting value into minHeap
func (m *minHeap) insert(value int) {
	m.array = append(m.array, value)
	m.heapifyBotToTop()
}

//Fn to maintain the minHeap property.
func (m *minHeap) heapifyBotToTop() {
	index := len(m.array) - 1
	for index > 0 {
		if m.array[parent(index)] > m.array[index] {
			m.array[parent(index)], m.array[index] = m.array[index], m.array[parent(index)]
		}
		index = parent(index)
	}
}

// To find the parent node of a node
func parent(index int) int {
	return (index - 1) / 2
}

//Extract the top node
func (m *minHeap) extract() {
	fmt.Print(m.array[0])
	fmt.Printf(" deleted! \n")
	m.array[0] = m.array[len(m.array)-1]
	m.array = m.array[:len(m.array)-1]
	m.heapifyTopToBot()
}

//Maintain the heap property from Top to Bottom
func (m *minHeap) heapifyTopToBot() {
	index := 0
	lc := lChild(index)
	rc := rChild(index)
	for index <= len(m.array)-1 {
		//edge case when there are only 2 elements in heap
		if len(m.array) == 2 {
			if m.array[lc] < m.array[index] {
				m.array[index], m.array[lc] = m.array[lc], m.array[index]
				index = lChild(index)
				return
			}
		}
		//normal case
		if m.array[lc] < m.array[index] && m.array[lc] < m.array[rc] {
			m.array[lc], m.array[index] = m.array[index], m.array[lc]
			index = lc
			fmt.Println("Hei")
		} else if m.array[rc] < m.array[index] && m.array[rc] < m.array[lc] {
			m.array[rc], m.array[index] = m.array[index], m.array[rc]
			index = rc
		} else {
			return
		}
	}
}

// To find the left Child of a node
func lChild(index int) int {
	return (index * 2) + 1
}

// To find the right Child of a node
func rChild(index int) int {
	return (index * 2) + 2
}

//MAIN fn
func main() {
	m := minHeap{}
	sampleArray := []int{7, 6, 5, 3, 4, 2, 1}
	for _, v := range sampleArray {
		m.insert(v)
	}
	fmt.Println(m.array)
	m.extract()
	fmt.Println(m.array)

}
