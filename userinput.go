package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
) 

func getInput(r *bufio.Reader, prompt string) (string, error ) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput(reader, "Create a new bill for: ")

	b := newBill(name)
	fmt.Println("Created new bill for:", b.name)

	return b
}

func promptOptions(b bill){

	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput(reader, "Choose and option (a - add item, s - save bill, t - add tip): ")
	
	switch option {
	case "a":
		name, _ := getInput(reader, "What's the item name")
		price, _ := getInput(reader, "What's the price?")

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)
		promptOptions(b)


		fmt.Println("Item Added ", name, p)
	case "s":
		fmt.Println("You chose to save the bill")
	case "t":
		tip, _ := getInput(reader, "Enter tip amount: ")
		fmt.Println(tip)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("tip added - ", t)
		promptOptions(b)
	default:
		fmt.Println("Invalid option")
		promptOptions(b)
	}
}

func main () {
	myBill := createBill()

	promptOptions(myBill)

	fmt.Println(myBill)
}