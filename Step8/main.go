package main

import (
	"fmt"
	"strings"
)

func main() {
	fn1, ln1 := getInitials("Nama User")
	fmt.Println(fn1, ln1)

	fn2, ln2 := getInitials("web CLoud")
	fmt.Println(fn2, ln2)

	fn3, ln3 := getInitials("Amazon")
	fmt.Println(fn3, ln3)
}

func getInitials(name string) (string, string) {
	var uname = strings.ToUpper(name)
	var s = strings.Split(uname, " ")

	var initials []string
	for _, v := range s {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}
