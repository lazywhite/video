package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u1 := User{
		Name: "alice",
		Age:  1,
	}
	u2 := User{
		Name: "bob",
		Age:  2,
	}

	users := []User{u1, u2}
	userRefs := []*User{&u1, &u2}
	structWithoutIndex(users, "alice")
	structWithIndex(users, "alice")

	refWithoutIndex(userRefs, "bob")
	refWithIndex(userRefs, "bob")
}

func structWithoutIndex(data []User, name string) {
	for _, u := range data {
		if u.Name == name {
			u.Age = 10
		}
	}
	fmt.Println(data)
}

func structWithIndex(data []User, name string) {
	for i, u := range data {
		if u.Name == name {
			data[i].Age = 10
		}
	}
	fmt.Println(data)
}

func refWithoutIndex(data []*User, name string) {
	for _, u := range data {
		if u.Name == name {
			u.Age = 20
		}
	}
	fmt.Println(*data[0], *data[1])
}

func refWithIndex(data []*User, name string) {
	for i, u := range data {
		if u.Name == name {
			data[i].Age = 10
		}
	}
	fmt.Println(*data[0], *data[1])
}
