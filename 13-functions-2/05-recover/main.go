package main

import (
	"fmt"
)

func main() {
	err := RunSafely(func() {
		Divide(10, 0)
	})
	fmt.Println(err)
}

func Divide(x, y float64) float64 {
	if y == 0 {
		panic("division by zero")
	}

	return x / y
}

func RunSafely(f func()) (err error) {

	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("Recovered panic: %v", r)
		}
	}()

	f()

	return
}
