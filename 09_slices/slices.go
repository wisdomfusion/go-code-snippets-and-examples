package main

import (
	"fmt"
	"strconv"
)

func main() {
	s1 := make([]string, 5, 10)
	for i := range s1 {
		s1[i] = "str" + strconv.Itoa(i)
	}
	fmt.Printf("s1=%v, len(s1)=%d, cap(s1)=%d", s1, len(s1), cap(s1))

	fmt.Println("s2")
	s2 := make([]string, 5)
	s2[0] = "a"
	s2[1] = "b"
	s2[2] = "c"
	s2[3] = "d"
	s2[4] = "e"

	fmt.Println(s2)
	fmt.Println(len(s2))

	s2 = append(s2, "f")
	fmt.Println(s2)
	fmt.Println(len(s2))

	c := make([]string, len(s2))
	copy(c, s2)
	fmt.Println(c)

	l := s2[:3]
	fmt.Println(l)
}
