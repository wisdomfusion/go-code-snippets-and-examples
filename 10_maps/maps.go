package main

import (
	"fmt"
	"sort"
)

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

	for k, v := range n {
		fmt.Println(k, v)
	}

	fmt.Printf("get not existed key from map: [%v]\n", n["none"])

	if name, ok := n["name"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key does not exists")
	}

	foo, ok := n["foo"]
	fmt.Println(foo, ok)
	delete(n, "foo")
	foo, ok = n["foo"]
	fmt.Println(foo, ok)

	fmt.Println("sort a map by key:")

	persons := map[string]int{"Gavin": 36, "Alice": 23, "Eve": 2, "Bob": 25}

	names := make([]string, 0, len(persons))
	for name := range persons {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Println(name, persons[name])
	}
}
