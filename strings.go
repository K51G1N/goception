package main

import "fmt"

func main() {

	age := 110
	name := "aang"
	element := "air"

	//Print
	fmt.Print("Hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	//Print line
	fmt.Println("greetings!")
	fmt.Println("goodbye!")
	fmt.Println("my name is:", name, "and my age is:", age)
	
	// Printf (formatted strings)
	fmt.Printf("My age is %d and my age is %s I am an %sbender and the avatar \n", age, name, element)
	fmt.Printf("age is of %T \n", age)
	fmt.Printf("You scored %0.2f points! \n", 123.123)

	//Sprintf saves string in variable for us
	var statement string = fmt.Sprintf("My age is %d and my age is %s I am an %sbender and the avatar \n", age, name, element)
	fmt.Println("The saved string:", statement)
}