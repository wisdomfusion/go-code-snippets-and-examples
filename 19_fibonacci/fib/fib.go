package fib

func Fibonacci() func() int {
	x, y := 0, 1

	return func() int {
		x, y = y, x+y
		return x
	}
}
