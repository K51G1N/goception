package main

import (
	"fmt"
	"strings"
)

func getInitials(name string) (firstname string, lastname string) {
	s := strings.ToUpper(name)
	names := strings.Split(s, " ")

	var initials []string

	for _, value := range names {
		initials = append(initials, value[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	firstName, secondName := getInitials("Michael Scott")
	fmt.Printf("Michael Scott your initials are: %v.%v \n", firstName, secondName)

	firstName1, secondName1 := getInitials("Dwight")
	fmt.Printf("Michael Scott your initials are: %v.%v \n", firstName1, secondName1)

	//doesn't work for no name entered. Will explore user input validation later on.
}
