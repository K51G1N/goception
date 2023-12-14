package main

import "fmt"

func main() {
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   7.99,
		"salad": 6.99,
		"tofu":  8.99,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// Add new item
	menu["cake"] = 5.99
	fmt.Println(menu)

	for key, value := range menu {
		fmt.Print(key, " : ", value, "\n")
	}

	// ints as key type
	phonebook := map[int]string{
		4567899: "Aang",
		1234508: "Katara",
		9876541: "Sokka",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[1234508])


	phonebook[1234508] = "Toph"
	fmt.Print(phonebook)
	fmt.Println(phonebook[1234508])

}
