package main

import (
	"fmt"
	"math"
)

func main() {
	var myByte byte = 'a'
	var myRune rune = '♥'

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)

	var r = 2.5
	var res = fmt.Sprintf("%.2f", math.Pi * r * r)
	fmt.Printf("res: %v", res)
}
