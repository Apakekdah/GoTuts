package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var greeting = "Hello there friends!"

	fmt.Println(strings.Contains(greeting, "Hello"))
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ll"))

	// original
	fmt.Println("Original value :", greeting)

	// substring
	fmt.Println("Substring:", greeting[6:11])

	// split
	fmt.Println("Split", strings.Split(greeting, " "))

	var ages = []int{56, 42, 13, 60, 27, 70, 28, 39, 5}
	fmt.Println(ages)

	sort.Ints(ages)

	fmt.Println("Sorted :", ages)

	fmt.Println("seachInt 54 from ages :", sort.SearchInts(ages, 54))

	names := [4]string{"Nama4", "Nama2", "Nama1", "Nama6"}
	fmt.Println(names)

	sort.Strings(names[:])
	fmt.Println("Sort :", names)

	fmt.Println("Search Wrong :", sort.SearchStrings(names[:], "nama2"))
	fmt.Println("Search Text :", sort.SearchStrings(names[:], "Nama2"))
}
