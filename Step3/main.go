package main

import "fmt"

func main() {
	//var ages [3]int = [3]int{25, 30, 35}
	//var ages = [3]int{25, 30, 35}
	ages := [3]int{25, 30, 35}

	names := [4]string{"Nama1", "Nama2", "Nama3", "Nama4"}
	names[1] = "Name2_Replace"

	fmt.Println(ages)
	fmt.Println(names)

	// slice
	var scores = []int{100, 50, 60}
	scores[2] = 25

	fmt.Println(scores)

	var newScores = append(scores, 30)
	fmt.Println(newScores)

	var rangeOne = names[1:3]
	var rangeTwo = names[2:]
	var rangeThree = names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	var newRangeOne = append(rangeOne, "appended")
	fmt.Println(newRangeOne)
}
