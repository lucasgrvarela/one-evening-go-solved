package main

type User struct {
	Name string
}

func NewUser(name string, storage Storage) User {
	user := User{Name: name}
	storage.Store(user)
	return User{}
}

type MapStorage struct {
	users map[string]User
}

func (m MapStorage) Store(user User) {
	m.users[user.Name] = user
}

type SliceStorage struct {
	users []User
}

func (s SliceStorage) Store(user User) {
	s.users = append(s.users, user)
}

func main() {
	mapStorage := MapStorage{users: make(map[string]User)}
	sliceStorage := SliceStorage{users: make([]User, 0)}

	NewUser("Jane", mapStorage)
	NewUser("Jack", sliceStorage)
}

type Storage interface {
	Store(user User)
}
