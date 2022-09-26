package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	name    string
	Address *Address
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"}}

	// deep copying
	jane := john
	jane.name = "Jane" //ok

	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country}

	john.Address.StreetAddress = "321 Baker St"

	fmt.Println(john, john.Address) //Duplicate address
	fmt.Println(jane, jane.Address) //Duplicate address
}
