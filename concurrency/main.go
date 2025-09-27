// Channel is use by GoRoutine to communicate.
// Cmd: go run main.go

package main

import (
	"fmt"
	"sync"
)

func main() {

	fmt.Println("Let's begin with Go Concurrency!!!")

	// here we made the channel count as 5
	myChn := make(chan int, 5)
	wg := &sync.WaitGroup{}

	// myChn <- 5
	// fmt.Println(myChn)
	wg.Add(2)

	// Receive ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanOpen := <-myChn

		fmt.Println(isChanOpen)
		fmt.Println(val)

		// fmt.Println(<-myChn)
		wg.Done()
	}(myChn, wg)

	// Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 0
		ch <- 5
		ch <- 6
		close(myChn)
		// ch <- 5
		// ch <- 6
		wg.Done()
	}(myChn, wg)

	wg.Wait()

}
