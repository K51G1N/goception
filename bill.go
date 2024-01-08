package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make a new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-15v ...$%v \n", k+":", v)
		total += v
	}
	
	// Add tip
	fs += fmt.Sprintf("%-15v ...$%0.2f \n", "tip:", b.tip)

	// total
	fs += fmt.Sprintf("%-15v ...$%0.2f \n", "total:", total+b.tip)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip //automatically dereferenced
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price

}
