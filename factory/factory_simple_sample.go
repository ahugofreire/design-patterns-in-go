package main

import "fmt"

// Interface que define o comportamento do objeto
type Animal interface {
	GetInfo() string
}

// Estrutura base que implementa a interface Animal
type BaseAnimal struct {
	Name string
	Age  int
}

// Método que retorna as informações básicas do animal
func (ba *BaseAnimal) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", ba.Name, ba.Age)
}

// Estrutura que representa um cão
type Dog struct {
	*BaseAnimal
	Breed string
}

// Método que retorna as informações do cão
func (d *Dog) GetInfo() string {
	return fmt.Sprintf("%s, Breed: %s", d.BaseAnimal.GetInfo(), d.Breed)
}

// Estrutura que representa um gato
type Cat struct {
	*BaseAnimal
	Color string
}

// Método que retorna as informações do gato
func (c *Cat) GetInfo() string {
	return fmt.Sprintf("%s, Color: %s", c.BaseAnimal.GetInfo(), c.Color)
}

// Factory que cria instâncias de Animal com base nos dados fornecidos
type AnimalFactory struct{}

func (af *AnimalFactory) CreateAnimal(animalType string, name string, age int, additionalData interface{}) Animal {
	switch animalType {
	case "dog":
		return &Dog{
			BaseAnimal: &BaseAnimal{
				Name: name,
				Age:  age,
			},
			Breed: additionalData.(string),
		}
	case "cat":
		return &Cat{
			BaseAnimal: &BaseAnimal{
				Name: name,
				Age:  age,
			},
			Color: additionalData.(string),
		}
	default:
		panic("Invalid animal type")
	}
}

func main() {
	// Uso do Factory para criar instâncias de animais
	animalFactory := &AnimalFactory{}

	dog := animalFactory.CreateAnimal("dog", "Rex", 3, "Labrador")
	cat := animalFactory.CreateAnimal("cat", "Fluffy", 2, "White")

	fmt.Println(dog.GetInfo()) // Name: Rex, Age: 3, Breed: Labrador
	fmt.Println(cat.GetInfo()) // Name: Fluffy, Age: 2, Color: White
}
