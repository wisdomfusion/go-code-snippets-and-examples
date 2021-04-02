package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("its before noon")
	default:
		fmt.Println("its after noon")
	}

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool value.")
		case int:
			fmt.Println("I'm a int value.")
		default:
			fmt.Printf("I've no idea about type %T\n", t)
		}
	}

	whoAmI(true)
	whoAmI(123)
	whoAmI("hey")
}
