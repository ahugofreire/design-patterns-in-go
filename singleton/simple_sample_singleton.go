package main

import (
	"fmt"
	"sync"
)

// Estrutura que representa o singleton
type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

// Função que retorna a única instância do singleton
func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			data: "Hello, world!",
		}
	})
	return instance
}

func main() {
	// Recupera a instância do singleton
	singleton := GetInstance()

	// Define um novo valor para a instância
	singleton.data = "Hello, singleton!"

	// Recupera a instância novamente
	singleton2 := GetInstance()

	// Imprime o valor da instância
	fmt.Println(singleton.data)          // Hello, singleton!
	fmt.Println(singleton2.data)         // Hello, singleton!
	fmt.Println(singleton == singleton2) // true
}
