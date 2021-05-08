package main

import "fmt"

type rectangle struct {
	width  int
	length int
}

func (r *rectangle) perimeter() int {
	return 2*r.width + 2*r.length
}

func (r *rectangle) area() int {
	return r.width * r.length
}

func main() {
	rect := rectangle{width: 3, length: 4}

	fmt.Println("Perimeter:", rect.perimeter())
	fmt.Println("Area:", rect.area())

	rectPointer := &rect
	fmt.Println("Perimeter:", rectPointer.perimeter())
	fmt.Println("Area:", rectPointer.area())
}
