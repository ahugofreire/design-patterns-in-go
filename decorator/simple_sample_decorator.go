package main

import "fmt"

// Component interface
type Component interface {
	Execute()
}

// ConcreteComponent struct
type ConcreteComponent struct{}

func (c *ConcreteComponent) Execute() {
	fmt.Println("Executando a operação do ConcreteComponent")
}

// Decorator struct
type Decorator struct {
	component Component
}

func (d *Decorator) Execute() {
	fmt.Println("Executando a operação do Decorator")
	d.component.Execute()
}

// ConcreteDecorator struct
type ConcreteDecorator struct {
	Decorator
}

func (cd *ConcreteDecorator) Execute() {
	fmt.Println("Executando a operação do ConcreteDecorator")
	cd.Decorator.Execute()
}

func main() {
	component := &ConcreteComponent{}

	decorator := &Decorator{
		component: component,
	}

	decoratedComponent := &ConcreteDecorator{
		Decorator: *decorator,
	}

	decoratedComponent.Execute()
}
