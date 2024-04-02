package main

import (
	"fmt"

	ioc "github.com/richhh7g/simple-ioc"
)

type User struct {
	Name  string
	Email string
}

func NewUser(name, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}

func (u *User) GetName() string {
	return u.Name
}

func main() {
	id := ioc.ContainerTag("CONTAINER_ID")
	container := ioc.NewContainer(&id)

	container.Register("User", NewUser("username", "email@domain.com"))

	user := container.Resolve("User").(*User)
	fmt.Println(user.GetName())
}
