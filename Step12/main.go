package main

import "fmt"

func main() {
	var myBill = newBill("Tagihan")

	fmt.Println(myBill)

	var format = myBill.format()

	fmt.Println(format)

	myBill.updateTip(5.56)

	fmt.Println(myBill.format())

	myBill.updateTipPointer(5.56)

	fmt.Println(myBill.format())

	for k := range myBill.items {
		delete(myBill.items, k)
	}

	myBill.addItem("Kentang", 3.12)  // kok bisa ??? hahaha
	myBill.addItem("Cumi-cumi", 1.5) // ini kok juga ??? hahaha

	fmt.Println(myBill.format())
}

// run with : go run main.go bill.go
