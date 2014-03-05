package main

import "fmt"

//START OMIT
func main() {
	fmt.Println("Printing a very large number first:")
	fmt.Println(aBigNumber())
	fmt.Println(aSmallNumber())
}

func aSmallNumber() int {
	return 7
}

func aBigNumber() int {
	return fib(42)
}

func fib(n int) int {
	if n <= 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)

}

//END OMIT
