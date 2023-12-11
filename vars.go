package main

import "fmt"

func vars () {
	
	fmt.Println("Hello World!")

	//strings
	var nameOne string = "Aang"
	var nameTwo = "Katara" // type inference
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	nameThree = "Sokka"
	println(nameThree)

	nameFour := "Toph" // no need for var and it uses type inference

	println(nameOne, nameTwo, nameThree, nameFour)

	//ints

	var  ageOne int = 10
	var ageTwo int = 12
	var ageThree = 14
	ageFour := 11

	println(ageOne, ageTwo, ageThree, ageFour)

	//bits and memory

	var smallBoi int8 = 127
	var sadBoi int8 = -128
	println(smallBoi, sadBoi)

	var biggerBoi uint = 255

	println(biggerBoi)

	//floats

	var daysToCommet float32 = 150
	println(daysToCommet)

	
}