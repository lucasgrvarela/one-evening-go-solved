package main

import (
	"errors"
)

var message = ""

func StoreMessage(m string) error {
	if m == "" {
		return errors.New("no message")
	}

	message = m

	return nil
}

func main() {
	MustStoreMessage("Hello!")
}

func MustStoreMessage(s string) {
	err := StoreMessage(s)
	if err != nil {
		panic(err)
	}
}
