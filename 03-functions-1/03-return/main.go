package main

import "fmt"

func main() {
	result := Sum(1, 2, 3, 4, 5)
	fmt.Println(result)
}

func Sum(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}
