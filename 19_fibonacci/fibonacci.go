package main

import (
	"fmt"
	"github.com/wisdomfusion/go_examples/19_fibonacci/fib"
)

func main() {
	f := fib.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
