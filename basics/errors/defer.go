package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wisdomfusion/go_examples/practices/fibonacci/fib"
)

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	pwd, _ := os.Getwd()
	fp := pwd + string(os.PathSeparator) + "fib.dat"
	writeFile(fp)
}
