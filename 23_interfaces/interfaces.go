package main

import (
	"fmt"
	"github.com/WisdomFusion/go_examples/23_interfaces/geometry"
)

type Geometry interface {
	Perimeter() float64
	Area() float64
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.Perimeter())
	fmt.Println(g.Area())
}

func main() {
	r := geometry.Rectangle{Width: 2, Length: 3}
	t := geometry.Triangle{B: 5, H: 3, A: 4, C: 6}
	c := geometry.Circle{Radius: 3}

	measure(r)
	measure(t)
	measure(c)
}
