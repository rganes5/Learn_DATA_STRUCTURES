package main

import "fmt"

func main() {
	sampleArray := []int{1, 2, 3, 4, 5}
	// var limit int
	// var sampleArray [10]int
	// fmt.Println("Enter the array limit:")
	// fmt.Scan(&limit)
	// fmt.Println("Enter the Array elements :")
	// for i := 0; i < limit; i++ {
	// 	fmt.Scan(&sampleArray[i])
	// }

	//COPY DECMONSTRATION

	newArray := make([]int, len(sampleArray))

	//COPY ELEMENTS TO ANOTHER ARRAY USING LOOPS

	for i, v := range sampleArray {
		newArray[i] = v
	}

	//COPY ELEMENTS TO ANOTHER ARRAY USING COPY FUNCTION

	// copy(newArray, sampleArray)

	//COPY ELEMENTS TO ANOTHER ARRAY USING APPEND FUNCTION

	// newArray = append(newArray, sampleArray...)

	fmt.Println("The sample array is: ")
	fmt.Printf("%v", sampleArray)
	fmt.Println("\nThe copied array is: ")
	fmt.Printf("%v", newArray)

}
