package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the number of terms: ")
	fmt.Scan(&n)

	a := 0
	b := 1
	c := 2
	Tib := []int{}

	for i := 0; i < n; i++ {
		Tib = append(Tib, a)
		// a = b
		// b = c						this is the usuall mistake
		// c = a + b + c

		a, b, c = b, c, a+b+c
	}

	fmt.Printf("Tib Series: %d", Tib)
}
