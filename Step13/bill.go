package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	var b = bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
func (b bill) format() string {
	return b.formatPointer()
}

func (b *bill) formatPointer() string {
	var fs = "Bill breakdown:\n"
	var total float64 = 0

	// list items
	if len(b.items) < 1 {
		fs += "No Items\n"
	} else {
		for k, v := range b.items {
			fs += fmt.Sprintf("%-25v ...%0.2f\n", k, v) // span-right 25 with (-)
			total += v
		}

		fs += fmt.Sprintf("%25v : %0.2f\n", "Total", total) // span-left 25
	}

	fs += fmt.Sprintf("Tip : %0.2f", b.tip)

	return fs
}

// update tip
func (b bill) updateTip(tip float64) {
	b.tip = tip
}
func (b *bill) updateTipPointer(tip float64) {
	//(*b).tip = tip <- bisa ini juga
	b.tip = tip
}

// add item
func (b bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b bill) save() {
	var data = []byte(b.format())

	if _, err := os.Stat("bill"); os.IsNotExist(err) {
		os.Mkdir("bill", 066)
	}

	// var err = os.WriteFile("bill/"+b.name+".txt", data, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	if err := os.WriteFile("bill/"+b.name+".txt", data, 0644); err != nil {
		panic(err)
	}
	fmt.Println("File has been saved")
}
