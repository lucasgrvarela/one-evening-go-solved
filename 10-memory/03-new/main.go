package main

import "fmt"

var counter int

func AllocateBuffer() *string {
	counter++
	if counter >= 4 {
		return nil
	}
	return new(string)
}

func main() {
	var buffers []*string
	for {
		b := AllocateBuffer()
		if b == nil {
			break
		}

		buffers = append(buffers, b)
	}

	fmt.Println("Allocated", len(buffers), "buffers")
}
