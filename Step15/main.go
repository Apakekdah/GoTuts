package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	var now = time.Now()
	var tgl = time.Date(2019, 10, 1, 0, 0, 0, 0, time.Local)

	fmt.Println(now.Format(time.ANSIC))

	var formatTgl = now.Format("2006-02-01 15:04:05")
	fmt.Println(formatTgl)

	var diff = now.Sub(tgl)
	fmt.Println(now, tgl)
	fmt.Println(diff)
	fmt.Println("Total hours :", diff.Hours())
	fmt.Println("Total minutes :", diff.Minutes())
	fmt.Println("Total seconds :", diff.Seconds())

	fmt.Println("Diff Days :", math.Floor(diff.Hours()/24))

	parseDt, err := time.Parse("2006-02-01 15:04:05", formatTgl)
	if err != nil {
		fmt.Println("Format error", err)
	} else {
		fmt.Println(parseDt)
	}
}
