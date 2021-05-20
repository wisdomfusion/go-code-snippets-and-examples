package main

import "fmt"

func main() {
	var arrApple [5]int
	fmt.Println(arrApple)

	arrApple[0] = 1
	fmt.Println(arrApple[0])

	fmt.Println(len(arrApple))

	arrBanana := [5]int{0:1, 3:4}
	fmt.Println(arrBanana)

	var arrTwoDimensions [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arrTwoDimensions[i][j] = i + j
		}
	}
	fmt.Println(arrTwoDimensions)
}
