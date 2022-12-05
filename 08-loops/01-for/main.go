package main

import "fmt"

func Alphabet(length int) (s []string) {
	for x := 0; x < length; x++ {
		s = append(s, characterByIndex(x))
	}
	return s
}

func main() {
	alphabet := Alphabet(26)
	fmt.Println(alphabet)
}

func characterByIndex(i int) string {
	return string(rune('a' + i))
}
