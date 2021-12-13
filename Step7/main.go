package main

import (
	"fmt"
	"math"
)

func main() {
	sayGreeting("Nama1")
	sayGreeting("Nama2")
	sayBye("Nama1")

	names := []string{"Nama1", "Nama2", "Nama3", "Nama4", "Name5"}

	cycleNames(names, sayGreeting)

	var a1 = circleArea(10.5)
	var a2 = circleArea(15)

	fmt.Printf("Area 1 %0.3f, Area 2 %0.3f", a1, a2)
}

func sayGreeting(name string) {
	fmt.Printf("Hello %v\n", name)
}

func sayBye(name string) {
	fmt.Printf("Bye %v\n", name)
}

func cycleNames(names []string, fn func(string)) {
	for _, value := range names {
		fn(value)
	}
}

func circleArea(radius float64) float64 {
	return math.Pi * radius * radius
}
