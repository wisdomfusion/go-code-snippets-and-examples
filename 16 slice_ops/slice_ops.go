package main

import "fmt"

func main() {
	s := make([]int, 10, 12)
	for i := range s {
		s[i] = i + 1
	}
	fmt.Println(s)

	s = append(s, []int{11, 12, 13, 14, 15}...)
	fmt.Println(s)

	fmt.Println("popping from start")
	fmt.Println(s[0])
	s = s[1:]
	fmt.Printf("s=%v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))

	fmt.Println("popping from end")
	fmt.Println(len(s))
	s = s[:len(s)-1]
	fmt.Printf("s=%v, len(s)=%d, cap(s)=%d\n", s, len(s), cap(s))

	fmt.Println("nil slice")
	var nilSlice []int
	fmt.Printf("nilSlice=%v, len(nilSlice)=%d, cap(nilSlice)=%d\n", nilSlice, len(nilSlice), cap(nilSlice))
	if nilSlice == nil {
		fmt.Println("nil slice!")
	}
}
