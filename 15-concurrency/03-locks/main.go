package main

import (
	"sync"
	"time"
)

var lock sync.Mutex

type User struct {
	Email string
}

type Storage struct {
	users map[string]User
}

func (s *Storage) AddUser(email string) {
	lock.Lock()
	defer lock.Unlock()
	_, ok := s.users[email]
	if !ok {
		s.users[email] = User{Email: email}
	}
}

func main() {
	storage := &Storage{users: make(map[string]User)}

	emails := []string{
		"alice@example.com",
		"kate@example.com",
		"joe@example.com",
		"rob@example.com",
		"patrick@example.com",
	}

	for _, email := range emails {
		go storage.AddUser(email)
	}

	time.Sleep(time.Millisecond * 100)
}
