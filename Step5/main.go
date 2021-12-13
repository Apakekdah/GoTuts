package main

import (
	"fmt"
)

func main() {
	var x = 0
	for x < 5 {
		fmt.Println("Value :", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("i Value :", i)
	}

	names := [4]string{"Nama1", "Nama2", "Nama3", "Nama4"}

	for x := 0; x < len(names); x++ {
		fmt.Println("Name :", names[x])
	}

	for index, v := range names {
		//fmt.Println("Pos At :", index, "Value :", v)
		fmt.Printf("Pos at %v with Value %v \n", index, v)
	}

	for _, value := range names {
		fmt.Printf("For Range Value : %v\n", value)
	}
}
