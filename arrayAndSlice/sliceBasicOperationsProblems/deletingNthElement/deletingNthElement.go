package main

import "fmt"

func main() {
	sampleSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("\nThe slice is: %v", sampleSlice)
	//DELETING LAST ELEMENT FROM SLICE
	sampleSlice = append(sampleSlice[:8], sampleSlice[9:]...)
	fmt.Printf("\nThe slice is: %v", sampleSlice)
}
