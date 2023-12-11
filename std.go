package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	greeting := "Hello there friends!"
	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hiii"))
	fmt.Println("original greeting: ", greeting)
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ie"))

	fmt.Println(strings.Split(greeting, " "))

	temperatures := []int{19, 22, 16, 33, 28, 24, 25, 8}

	sort.Ints(temperatures) //alters original slice
	fmt.Println(temperatures)

	index := sort.SearchInts(temperatures, 28)
	fmt.Println((index)) // returns length, if not found returns len(slice)+1

	teamAvatar := []string{"Aang", "Katara", "Sokka", "Toph", "Zuko", "Appa", "Momo"}

	sort.Strings(teamAvatar)
	fmt.Println(teamAvatar)

	fmt.Println(sort.SearchStrings(teamAvatar, "Katara"))
}
