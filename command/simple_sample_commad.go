package main

import "fmt"

// Interface Command que define o método Execute()
type Command interface {
	Execute()
}

// Struct para o comando de impressão
type PrintCommand struct {
	Message string
}

// Método Execute() para o comando de impressão
func (p *PrintCommand) Execute() {
	fmt.Println(p.Message)
}

// Struct para o comando de adição
type AddCommand struct {
	Value int
}

// Método Execute() para o comando de adição
func (a *AddCommand) Execute() {
	fmt.Printf("Add %d to the current value\n", a.Value)
}

// Struct para o comando de subtração
type SubCommand struct {
	Value int
}

// Método Execute() para o comando de subtração
func (s *SubCommand) Execute() {
	fmt.Printf("Subtract %d from the current value\n", s.Value)
}

// Struct Invoker que gerencia a fila de comandos
type Invoker struct {
	commands []Command
}

// Método para adicionar comandos à fila
func (i *Invoker) AddCommand(c Command) {
	i.commands = append(i.commands, c)
}

// Método para executar todos os comandos na fila
func (i *Invoker) ExecuteCommands() {
	for _, c := range i.commands {
		c.Execute()
	}
}

func main() {
	// Cria um invoker
	invoker := Invoker{}

	// Adiciona alguns comandos à fila
	invoker.AddCommand(&PrintCommand{Message: "Hello, world!"})
	invoker.AddCommand(&AddCommand{Value: 10})
	invoker.AddCommand(&SubCommand{Value: 5})

	// Executa todos os comandos na fila
	invoker.ExecuteCommands()
}

// Nesse exemplo, temos a interface Command que define o método Execute(), que é implementado por cada comando específico (no caso, PrintCommand, AddCommand e SubCommand). Cada comando encapsula
// uma solicitação específica e implementa o método Execute() para executar essa solicitação.
// Também temos a struct Invoker, que gerencia a fila de comandos e implementa os métodos AddCommand() para adicionar comandos à fila e ExecuteCommands() para executar todos os comandos na fila.
// No exemplo acima, criamos um invoker, adicionamos alguns comandos à fila e executamos todos os comandos na fila.
// Dessa forma, o padrão Command permite que você encapsule uma solicitação como um objeto, permitindo que você faça mais com essa solicitação posteriormente.
//Por exemplo, você pode adicionar comandos a uma fila, desfazer comandos, fazer log de comandos e assim por diante. Isso torna o código mais modular e mais fácil de manter e evoluir no futuro
