package main

import (
	"fmt"
	"strings"
)

func main() {
	myStr := " helLo world !"
	name := "Ganesh"
	fmt.Println(myStr)
	for index, v := range myStr {
		fmt.Printf("Index number of %c is %d and byte is %v \n", v, index, v)
	}
	//Length
	fmt.Println(len(myStr))

	//Accessing charecters
	fmt.Printf("Charecter representation : %c\n", myStr[1])

	//Concatenate
	fmt.Print(myStr + "" + name)
	fmt.Println("")

	//Calling functions
	basicCasing(myStr)
	trim(myStr)
	replace(myStr)

}

//BASIC MANIPULATIONS

func basicCasing(myStr string) {

	//Uppercase
	fmt.Println(strings.ToUpper(myStr))

	//Lowercase
	fmt.Println(strings.ToLower(myStr))

	//Titlecase
	fmt.Println(strings.Title(myStr))
}

//TRIMMING

func trim(myStr string) {
	fmt.Print("Trim Right :", strings.TrimRight(myStr, " "))
	fmt.Println(" ")
	fmt.Print("Trim Right :", strings.TrimLeft(myStr, " "))
	fmt.Println(" ")
}

//REPLACING

func replace(myStr string) {
	fmt.Println(strings.Replace(myStr, "world", "India", 5))
	fmt.Println(strings.ReplaceAll(myStr, "world", "Keralam"))
}
