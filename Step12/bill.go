package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	var b = bill{
		name:  name,
		items: map[string]float64{"Pie": 5.46, "Cake": 12.52},
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

func (b *bill) addItemPointer(name string, price float64) {
	b.items[name] = price
}
