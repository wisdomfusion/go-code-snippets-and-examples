package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err)
		} else {
			panic(fmt.Sprintf("i don't know what to do: %v", r))
		}
	}()
	//panic(errors.New("this is an error"))

	//a := 0
	//b := 5 / a
	//fmt.Println(b)
	panic(111)
}

func main() {
	tryRecover()
}
