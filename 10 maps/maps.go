package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["key1"] = 1
	m["key2"] = 2

	fmt.Println(m)

	v1 := m["key1"]
	fmt.Println(v1)

	delete(m, "key2")
	fmt.Println(m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}
