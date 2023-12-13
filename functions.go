package main

import (
	"fmt"
	"math"
)

func sayGreeting(name string) {
	fmt.Printf("Greetings %v! \n", name)
}

func sayFarewell(name string) {
	fmt.Printf("Farewell Dear %v :( \n", name)
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func AreaOfCircle(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	sayGreeting("Jim")
	sayFarewell("Michael")
	cycleNames([]string{"Appa", "Momo", "Oogi"}, sayGreeting)
	cycleNames([]string{"Appa", "Momo", "Oogi"}, sayFarewell)
	a1 := AreaOfCircle(12.12)
	a2 := AreaOfCircle(11.11)
	a3 := AreaOfCircle(13.13)

	fmt.Printf("Circle 1: %0.2f , Circle 2: %0.2f, Circle 3: %0.2f", a1, a2, a3)
}
