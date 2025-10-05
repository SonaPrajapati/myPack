// Sets max CPU threads and logs identifier for runtime configuration.
// Parallelism in go.

package main

import (
	"fmt"
	"runtime"
)

func configureRuntime(text string) {
	fmt.Println("this is the given text: ", text)
	runtime.GOMAXPROCS(2)
}

func main() {
	configureRuntime("Example Text")
	fmt.Println("Hi team")
}
