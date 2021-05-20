package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [10]int

	for i := 0; i < 10; i++ {
		go func(ii int) {
			for {
				a[ii]++
				runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println(a)
}
