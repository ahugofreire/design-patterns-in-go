package main

import "fmt"

type Person struct {
	//address
	StreetAddress, Postcode, City string

	//job
	CompanyName, Position string
	AnnualIncome          int
}

type PersonBuilder struct {
	person *Person
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*b}
}
func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{*b}
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}

type PersonAddressBuilder struct {
	PersonBuilder
}

type PersonJobBuilder struct {
	PersonBuilder
}

func (it *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	it.person.StreetAddress = streetAddress
	return it
}

func (it *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	it.person.City = city
	return it
}

func (it *PersonAddressBuilder) WithPostcode(postcode string) *PersonAddressBuilder {
	it.person.Postcode = postcode
	return it
}

func (pjb *PersonJobBuilder) At(companyName string) *PersonJobBuilder {
	pjb.person.CompanyName = companyName
	return pjb
}

func (pjb *PersonJobBuilder) AsA(position string) *PersonJobBuilder {
	pjb.person.Position = position
	return pjb
}

func (pjb *PersonJobBuilder) Earning(annualIncome int) *PersonJobBuilder {
	pjb.person.AnnualIncome = annualIncome
	return pjb
}

func (b *PersonBuilder) Build() *Person {
	return b.person
}

func main() {
	pb := NewPersonBuilder()
	pb.Lives().
		At("123 London Road").
		In("London").
		WithPostcode("SW12CD").
		Works().
		At("Fabrikam").
		AsA("Programmer").
		Earning(122111)
	person := pb.Build()
	fmt.Println(person)
}
