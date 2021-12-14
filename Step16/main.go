// Sample from youtube : https://www.youtube.com/watch?v=LvgVSSpwND8

package main

import (
	"fmt"
	"time"
)

func main() {
	/// -> 1
	//go count("sheep")
	//go count("fish")
	//time.Sleep(time.Second * 3)
	//fmt.Scanln()

	/// -> 2
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	count("sheep")
	// 	wg.Done()
	// }()
	// wg.Wait()

	/// -> 3
	// var c = make(chan string)
	// go countChannel("sheep", c)
	// //var msg = <-c
	// //fmt.Println(msg) // will terminate
	// // 1
	// // for {
	// // 	msg, open := <-c
	// // 	if !open {
	// // 		break
	// // 	}
	// // 	fmt.Println(msg)
	// // }
	// // 2 simplify
	// for msg := range c {
	// 	fmt.Println(msg)
	// }

	/// -> 4
	// var c = make(chan string, 2)
	// c <- "Hello"
	// c <- "World"
	// c <- "Yeah" // make app deadlock because no listener for chan
	// var msg = <-c
	// fmt.Println(msg)
	// msg = <-c
	// fmt.Println(msg)
	// // deadlock in here, no listener for chan

	// /// -> 5
	// var c1 = make(chan string)
	// var c2 = make(chan string)
	// go func() {
	// 	for {
	// 		c1 <- "Every 500ms"
	// 		time.Sleep(time.Millisecond * 500)
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		c2 <- "Every 2 seconds"
	// 		time.Sleep(time.Second * 2)
	// 	}
	// }()
	// // 1 locking
	// // for {
	// // 	fmt.Println(<-c1)
	// // 	fmt.Println(<-c2)
	// // }
	// // 2 concurrent
	// for {
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println(msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println(msg2)
	// 	}
	// }

	/// 6
	// var jobs = make(chan int, 100)
	// var results = make(chan int, 100)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// go worker(jobs, results)
	// for i := 0; i < 100; i++ {
	// 	jobs <- i
	// }
	// close(jobs)
	// for i := 0; i < 100; i++ {
	// 	fmt.Println("Fib :", <-results)
	// }
	// close(results)

	fmt.Println("App End!")
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}

func countChannel(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c) // stop chan
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func worker(job <-chan int, result chan<- int) {
	for n := range job {
		result <- fib(n)
	}
}
