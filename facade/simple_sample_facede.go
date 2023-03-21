package main

import "fmt"

// Subsistema complexo
type CPU struct{}

func (c *CPU) freeze() {
	fmt.Println("CPU: freeze")
}

func (c *CPU) jump(position int) {
	fmt.Printf("CPU: jump to %d\n", position)
}

func (c *CPU) execute() {
	fmt.Println("CPU: execute")
}

// Subsistema complexo
type Memory struct{}

func (m *Memory) load(position int, data []byte) {
	fmt.Printf("Memory: load data %s at position %d\n", string(data), position)
}

// Subsistema complexo
type HardDrive struct{}

func (h *HardDrive) read(position int, size int) []byte {
	fmt.Printf("HardDrive: read data at position %d, size %d\n", position, size)
	return make([]byte, size)
}

// Facade
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

// Método que inicia o sistema
func (c *ComputerFacade) Start() {
	fmt.Println("Computer: start")
	c.cpu.freeze()
	c.memory.load(0, c.hardDrive.read(0, 1024))
	c.cpu.jump(0)
	c.cpu.execute()
}

func main() {
	computer := NewComputerFacade()
	computer.Start()
}
