package main

import "fmt"

func main() {
	var menu = map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// loops map
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// ints as key type
	var phonebook = map[int]string{
		123: "Nama 1",
		456: "Nama 2",
		789: "Nama 3",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[456])

	phonebook[137] = "Gila"

	fmt.Println(phonebook)

	phonebook[456] = "Nama 2 Dong"
	phonebook[789] = "Nama 3 Dong"

	fmt.Println(phonebook)
}
