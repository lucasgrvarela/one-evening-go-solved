package main

func main() {
	_ = Sum(1, 2, 3, 4, 5)
}

func Sum(i ...int) (sum int) {
	for _, v := range i {
		sum += v
	}
	return sum
}
