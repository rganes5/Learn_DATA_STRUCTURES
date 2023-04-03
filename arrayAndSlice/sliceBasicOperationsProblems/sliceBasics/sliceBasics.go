package main

import (
	"fmt"
)

func main() {
	// var mySlice int[]
	slice := []int{1, 2, 3, 4}
	slice = append(slice, 5, 6, 7, 8, 9)
	fmt.Println(slice)
	fmt.Println("...........................................................................")

	var my_slice_1 = make([]int, 4, 7)
	fmt.Printf("Slice 1 = %v, \nlength = %d, \ncapacity = %d\n",
		my_slice_1, len(my_slice_1), cap(my_slice_1))

	newSlice := []string{"Hello", "world"}
	fmt.Println(newSlice)
	newSlice = append(newSlice, "Yes")

	for i, v := range newSlice {
		fmt.Println(i, v)
	}

	fmt.Println("...........................................................................")

	//CREATING A MULTIDIMENTIONAL SLICE

	multiSlice := [][]string{
		{"Hello", "world"},
		{"Hello", "world"},
		{"Hello", "world"},
	}
	multiSlice2 := [][]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(multiSlice)
	fmt.Println(multiSlice2)
}
