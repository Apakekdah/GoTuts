package main

import "fmt"

var score = 99.5

func main() {
	sayHello("Woi")

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()
}

// must run : go run main.go greetings.go
