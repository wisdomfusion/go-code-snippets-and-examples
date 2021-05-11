package main

import (
	"fmt"
	"github.com/WisdomFusion/go_examples/19_fibonacci/fib"
)

func main() {
	f := fib.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
