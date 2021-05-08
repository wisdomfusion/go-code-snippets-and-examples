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

	dayOfWeek := 6
	switch dayOfWeek {
	case 1: fmt.Println("Monday")
	case 2: fmt.Println("Tuesday")
	case 3: fmt.Println("Wednesday")
	case 4: fmt.Println("Thursday")
	case 5: fmt.Println("Friday")
	case 6: {
		fmt.Println("Saturday")
		fmt.Println("Weekend")
	}
	case 7: {
		fmt.Println("Sunday")
		fmt.Println("Weekend")
	}
	default: fmt.Println("Invalid day")
	}

	switch dayOfWeek := 6; dayOfWeek {
	case 1: fmt.Println("Monday")
	case 2: fmt.Println("Tuesday")
	case 3: fmt.Println("Wednesday")
	case 4: fmt.Println("Thursday")
	case 5: fmt.Println("Friday")
	case 6: {
		fmt.Println("Saturday")
		fmt.Println("Weekend")
	}
	case 7: {
		fmt.Println("Sunday")
		fmt.Println("Weekend")
	}
	default: fmt.Println("Invalid day")
	}

	switch dayOfWeek := 5; dayOfWeek {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid Day")
	}

	BMI := 21.0
	switch {
	case BMI < 18.5:
		fmt.Println("You're underweight")
	case BMI >= 18.5 && BMI < 25.0:
		fmt.Println("Your weight is normal")
	case BMI >= 25.0 && BMI < 30.0:
		fmt.Println("You're overweight")
	default:
		fmt.Println("You're obese")
	}
}
