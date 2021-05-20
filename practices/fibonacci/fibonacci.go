package main

import (
	"fmt"

	"github.com/wisdomfusion/go_examples/practices/fibonacci/fib"
)

func main() {
	f := fib.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
