package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, rdr *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := rdr.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	var rdr = bufio.NewReader(os.Stdin)

	fmt.Print("Create new bill name : ")
	name, _ := rdr.ReadString('\n')
	name = strings.TrimSpace(name)

	var bill = newBill(name)

	fmt.Println("Create bill :", name)

	return bill
}

func createBillRefactor() bill {
	var rdr = bufio.NewReader(os.Stdin)

	name, _ := getInput("Create new bill name : ", rdr)

	var bill = newBill(name)

	fmt.Println("Create bill :", name)

	return bill
}

func promptOption(bill bill) {
	var rdr = bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose options (a = add item, t = add tipe, s = save bill) : ", rdr)
	//fmt.Println(opt)

	switch opt {
	case "a":
		//fmt.Println("You choose a")

		name, _ := getInput("Item name : ", rdr)
		price, _ := getInput("Price : ", rdr)
		fmt.Println(name, price)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Printf("Failed to parse price %q into number\n", price)
			promptOption(bill)
			return
		}

		bill.addItem(name, p)

		fmt.Printf("Added item %q with price %0.2f\n", name, p)
		promptOption(bill)
	case "t":
		//fmt.Println("You choose t")

		tip, _ := getInput("Add Tip : ", rdr)
		fmt.Println(tip)

		t, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Printf("Failed to parse tip %q into number\n", tip)
			promptOption(bill)
			return
		}

		bill.updateTipPointer(t)

		fmt.Printf("Update tip %0.2f\n", t)
		promptOption(bill)
	case "s":
		//fmt.Println("You choose s")

		bill.save()

		fmt.Println("File has been saved -", bill.name)
	default:
		fmt.Printf("Unknown option %q\n", opt)
		promptOption(bill)
	}
}

func main() {
	var myBill = createBillRefactor()

	promptOption(myBill)

	fmt.Println(myBill)
}

// run with : go run main.go bill.go
