package main

import "fmt"

func WordGenerator(words []string) func() string {
	i := 0
	return func() string {
		word := words[i]

		i++
		if i >= len(words) {
			i = 0
		}

		return word
	}
}

func main() {
	continents := []string{
		"Africa",
		"Antarctica",
		"Asia",
		"Australia",
		"Europe",
		"North America",
		"South America",
	}

	generator := WordGenerator(continents)

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
