package main

import "fmt"

func main() {
	s := make([]string, 5)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s[3] = "d"
	s[4] = "e"

	fmt.Println(s)
	fmt.Println(len(s))

	s = append(s, "f")
	fmt.Println(s)
	fmt.Println(len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[:3]
	fmt.Println(l)
}
