package main

import "fmt"

var products = map[int]string{
	1: "Book",
	2: "Video Course",
	3: "Lecture",
	4: "Talk",
	5: "Training",
}

func main() {
	ids := Keys(products)
	names := Values(products)

	fmt.Println("Prouct IDs:", ids)
	fmt.Println("Product names:", names)
}

func Keys(m map[int]string) (i []int) {
	for k := range m {
		i = append(i, k)
	}
	return i
}

func Values(m map[int]string) (s []string) {
	for _, v := range m {
		s = append(s, v)
	}
	return s
}
