package main

import "fmt"

func main() {
	myBill := newBill("Gokus's bill")

	myBill.addItem("onion rings", 4.50)
	myBill.addItem("pepperoni pizza", 8.95)
	myBill.addItem("malva pudding", 3.95)


	myBill.updateTip(17.99)

	fmt.Println(myBill.format())


}