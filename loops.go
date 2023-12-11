package main

import "fmt"

func main() {
	x := 0

	for x < 5 {
		fmt.Println("the value of x is: ", x)
		x++
	}

	for i := 0; i < 5; i++{
		fmt.Println("the value of i is: ", i)
	}
	
	theOfficeCharacters := []string{"Jim", "Pam", "Dwight", "Michael"}

	for i := 0; i < len(theOfficeCharacters); i++ {

		fmt.Println("Hi ", theOfficeCharacters[i])
	}

	for index, character := range(theOfficeCharacters) {
		index++
		fmt.Printf("Hi %s you're employee %d \n", character, index)
	}

	fmt.Println()
}