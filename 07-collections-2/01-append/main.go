package main

import "fmt"

var users []string

func main() {
	AddUser("Alice")
	AddUser("Bob")
	fmt.Println(users)
}

func AddUser(s string) {
	users = append(users, s)
}
