package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func values() (int, int) {
	return 1, 2
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	i, j := values()
	fmt.Println(i)
	fmt.Println(j)

	_, k := values()
	fmt.Println(k)
}
