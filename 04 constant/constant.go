package main

import (
	"fmt"
	"math"
)

const str string = "04 constant"

func main() {
	fmt.Println(str)

	const n = 50000000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}