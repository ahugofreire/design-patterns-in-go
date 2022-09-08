package main

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

// ISP
type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct {
}

func (m MyPrinter) Print(d Document) {

}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {

}

func (p Photocopier) Scan(d Document) {

}

type MultiFunctionDevice interface {
	Printer
	Scanner
	//Fax
}

// decorator
type MultiFunctionMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFunctionMachine) Print(d Document) {
	m.printer.Print(d)
}

func main() {

}
