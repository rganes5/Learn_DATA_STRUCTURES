package main

import "fmt"

func main() {
	sampleSlice := []int{2, 5, 0, 8, 9}
	fmt.Printf("The original array : %v", sampleSlice)
	max := sampleSlice[0]
	for i := 1; i < 5; i++ {
		if max > sampleSlice[i] {
			temp := max
			max = sampleSlice[i]
			sampleSlice[i] = temp
		}
	}
	fmt.Printf("\nThe smallest element is: %v", max)

}
