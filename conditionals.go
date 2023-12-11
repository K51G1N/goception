package main

import "fmt"


func main() {
	age := 75
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 75)
	fmt.Println(age != 60)

	if age < 100 {
		fmt.Println("age is less than 100")
	} else if age < 60 {
		fmt.Println("age is less than 60")
	} else {
		fmt.Println("Age is not less than 50")
	}

	teamAvatar := []string{"Aang", "Katara", "Sokka", "Toph", "Zuko", "Appa", "Momo"}

	for index, name := range(teamAvatar) {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breaking Pos", index)
			break
		}

		fmt.Printf(("The value at pos %v is %v \n"), index, name)
	}
}