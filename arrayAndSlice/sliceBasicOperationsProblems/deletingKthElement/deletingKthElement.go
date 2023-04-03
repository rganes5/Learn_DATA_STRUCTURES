package main

import "fmt"

func main() {
	sampleSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("\nThe slice is: %v", sampleSlice)
	//DELETING 3 FROM SLICE
	sampleSlice = append(sampleSlice[:2], sampleSlice[3:]...)
	fmt.Printf("\nThe slice is: %v", sampleSlice)
}
