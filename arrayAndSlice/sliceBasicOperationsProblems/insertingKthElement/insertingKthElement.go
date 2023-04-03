package main

import "fmt"

func main() {
	sampleSlice := []int{1, 2, 3, 4, 6, 7, 8, 9}
	fmt.Printf("The original array : %v", sampleSlice)

	//inserting at 4th index
	value := 5
	sampleSlice = append(sampleSlice, 0)
	fmt.Printf("\nAfter inserting a free space : %v", sampleSlice)

	copy(sampleSlice[4:], sampleSlice[3:])

	fmt.Println(sampleSlice)
	sampleSlice[2] = value
	fmt.Printf("\nAfter insertion : %v", sampleSlice)

}
