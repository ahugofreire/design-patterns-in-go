package main

import "fmt"

// Interface que define o comportamento do prototype
type Prototype interface {
	Clone() Prototype
}

// Estrutura que representa um usuário
type User struct {
	Name     string
	Age      int
	Location string
}

// Método que cria um clone do usuário
func (u *User) Clone() Prototype {
	return &User{
		Name:     u.Name,
		Age:      u.Age,
		Location: u.Location,
	}
}

func main() {
	// Criação do usuário original
	user1 := &User{
		Name:     "John",
		Age:      30,
		Location: "New York",
	}

	// Clone do usuário original
	user2 := user1.Clone().(*User)

	fmt.Println(user1) // &{John 30 New York}
	fmt.Println(user2) // &{John 30 New York}

	// Mudança da localização do usuário clone
	user2.Location = "London"

	fmt.Println(user1) // &{John 30 New York}
	fmt.Println(user2) // &{John 30 London}
}
