package main

import "fmt"

func updateName(name string) {
	name = "wedge"
}

func updateNameByPointer(name *string) {
	*name = "wedge"
}

func updateMenu(menu map[string]float64) {
	menu["coffee"] = 13.95
}

func main() {
	// group types by values -> string, ints, floats, bools, arrays, structs
	var name = "nama"

	updateName(name)

	fmt.Println(name)

	// pointer
	fmt.Println("Pointer :", &name)

	var m = &name

	fmt.Println("Memory address :", m)
	fmt.Println("Value at Memory :", *m)

	updateNameByPointer(m)

	fmt.Println("After update name :", *m)

	fmt.Println("Name begin :", name)

	// group types by reference -> slices, maps, functions
	var menu = map[string]float64{
		"pie":       5.95,
		"ice cream": 4.99,
	}

	fmt.Println("First menu :", menu)

	updateMenu(menu)

	fmt.Println("After update :", menu)
}
