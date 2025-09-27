// Here wee have used WaitGroup, Mutex also gone through RWMutex.
// WaitGroup: This is basically used to ensure every goroutine should be execute and none of goroutine left unused.
// Mutex: This is used to lock the goroutine while execution to prevent the intervention and intruption by any another goroutine.
// RWMutex: This is ReadWriteMutex this ensure Lock for read and write operation, this helps to prevent mis-handling of read and write operation and secure the workflow of the goroutine.
// RWMutex: this uses RLock & RUnlock
// Cmd: go run --race waitGroup.go

package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("The initial statement: ")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []string{"Zero"}

	wg.Add(3)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		mut.Lock()
		fmt.Println("The first go routine")
		score = append(score, "first")
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		mut.Lock()
		fmt.Println("The second go routine")
		score = append(score, "second")
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		mut.Lock()
		fmt.Println("The third go routine")
		score = append(score, "third")
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	// defer wg.Done()
	fmt.Println(score)

}
