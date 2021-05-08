package main

import "fmt"

func swap(a, b int) (int, int) {
	return b, a
}

func swapByPointer(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	fmt.Println(swap(1, 2))

	a, b := 3, 4
	swapByPointer(&a, &b)
	fmt.Println(a, b)
}
