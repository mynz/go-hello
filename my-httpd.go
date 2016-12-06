package main

import "fmt"

// import "log"

func fib() func() int {
	a, b := 0, 0
	// a, b := 0, 1
	// a, b := 1, 0

	return func() int {

		ret := a + b

		a, b = ret, a

		if b == 0 {
			b = 1
		}

		return ret
	}
}

func main() {

	fn := fib()
	for i := 0; i < 5; i++ {
		fmt.Println("fib: ", fn())
	}

}
