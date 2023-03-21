package main

import "fmt"

// Define a struct para o objeto a ser construído
type Pizza struct {
	dough   string
	sauce   string
	topping string
}

// Define a interface para o Builder
type PizzaBuilder interface {
	AddDough() PizzaBuilder
	AddSauce() PizzaBuilder
	AddTopping() PizzaBuilder
	Build() Pizza
}

// Define a struct para o Builder
type HawaiianPizzaBuilder struct {
	pizza Pizza
}

// Implementa os métodos da interface PizzaBuilder
func (b *HawaiianPizzaBuilder) AddDough() PizzaBuilder {
	b.pizza.dough = "regular"
	return b
}

func (b *HawaiianPizzaBuilder) AddSauce() PizzaBuilder {
	b.pizza.sauce = "tomato"
	return b
}

func (b *HawaiianPizzaBuilder) AddTopping() PizzaBuilder {
	b.pizza.topping = "ham and pineapple"
	return b
}

func (b *HawaiianPizzaBuilder) Build() Pizza {
	return b.pizza
}

// Define a struct para o Director
type PizzaDirector struct {
	builder PizzaBuilder
}

// Define um método para alterar o Builder utilizado
func (d *PizzaDirector) SetBuilder(builder PizzaBuilder) {
	d.builder = builder
}

// Define um método para construir o objeto
func (d *PizzaDirector) BuildPizza() Pizza {
	return d.builder.AddDough().AddSauce().AddTopping().Build()
}

func main() {
	// Utiliza o Builder para construir um objeto
	hawaiianBuilder := &HawaiianPizzaBuilder{}
	director := &PizzaDirector{builder: hawaiianBuilder}
	pizza := director.BuildPizza()
	fmt.Printf("Dough: %s\n", pizza.dough)
	fmt.Printf("Sauce: %s\n", pizza.sauce)
	fmt.Printf("Topping: %s\n", pizza.topping)
}
