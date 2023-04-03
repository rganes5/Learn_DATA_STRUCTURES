package main

import "fmt"

func main() {
	sampleSlice := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("The original array : %v", sampleSlice)

	value := 1
	sampleSlice = append(sampleSlice, 0)
	fmt.Printf("\nAfter inserting a free space : %v", sampleSlice)

	copy(sampleSlice[1:], sampleSlice[0:])
	sampleSlice[0] = value
	fmt.Printf("\nAfter insertion : %v", sampleSlice)
}
