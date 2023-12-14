package main

import "fmt"

func main() {
	myBill := newBill("Gokus's bill")

	fmt.Println(myBill.format())


}