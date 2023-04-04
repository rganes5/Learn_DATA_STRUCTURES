package main

import "fmt"

type maxHeap struct {
	array []int
}

func heapSort(array []int) {
	n := len(array)

	//To make the array a maxHeap.
	for i := (len(array) / 2) - 1; i >= 0; i-- {
		heapify(array, n, i)
	}
	//loop to make the array sorted
	// for i := n - 1; i >= 0; i-- {
	// 	array[0], array[i] = array[i], array[0]
	// 	heapify(array, i, 0)
	// }
}

func delete(array []int) []int {
	array[0] = array[len(array)-1]
	array = array[:len(array)-1]
	heapify(array, len(array), 0)
	return array
}

func heapify(array []int, n int, i int) {
	largest := i
	lc := lChild(i)
	rc := rChild(i)
	if lc < n && array[largest] < array[lc] {
		largest = lc
	}
	if rc < n && array[largest] < array[rc] {
		largest = rc
	}
	if largest != i {
		array[i], array[largest] = array[largest], array[i]
		heapify(array, n, largest)
	}
}

func lChild(index int) int {
	return (index * 2) + 1
}
func rChild(index int) int {
	return (index * 2) + 2
}
func main() {
	sampleArray := []int{10, 20, 15, 12, 40, 25, 18}
	heapSort(sampleArray)
	fmt.Println(sampleArray)
	sampleArray = delete(sampleArray)
	fmt.Println("After deletion")
	fmt.Println(sampleArray)
}
