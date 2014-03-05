package main

import "fmt"

//START OMIT
var results = make(chan int) // HL

func main() {
	go aBigNumber()
	go aSmallNumber()

	fmt.Println("Printing in whatever order they become available:")
	fmt.Println(<-results) // HL
	fmt.Println(<-results)
}

func aSmallNumber() {
	results <- 7 // HL
}

func aBigNumber() {
	results <- fib(42) // HL
}

//END OMIT

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
