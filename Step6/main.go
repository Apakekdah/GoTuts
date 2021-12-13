package main

import "fmt"

func main() {
	var age = 45

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Less than 30")
	} else if age < 40 {
		fmt.Println("Less than 40")
	} else {
		fmt.Println("Greater than 40")
	}

	names := []string{"Nama1", "Nama2", "Nama3", "Nama4", "Name5"}

	for index, val := range names {
		if index == 1 {
			fmt.Println("Continued at", index)
			continue
		} else if index > 2 {
			fmt.Println("Break at ", index)
			break
		}
		fmt.Printf("Value at pos %v is %v \n", index, val)
	}
}
