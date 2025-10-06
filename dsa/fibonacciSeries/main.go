package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Enter number of terms: ")
	fmt.Scan(&n)

	a, b := 0, 1

	Fib := []int{}

	// fmt.Println("Fibonacci Series:")

	for i := 0; i < n; i++ {
		Fib = append(Fib, i)
		a, b = b, a+b
	}
	fmt.Printf("Fibonacci Series: %d", Fib)

}
