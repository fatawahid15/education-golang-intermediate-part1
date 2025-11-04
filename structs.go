package main

import (
	"fmt"
)

type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address
	PhoneHomeCell // <= embedded anonymous field

}

type Address struct {
	city    string
	country string
}

type PhoneHomeCell struct {
	home string
	cell string
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *Person) incrementAgeByOne() { // <= Pointer equivalent shi
	p.age++
}

func main() {

	p := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "London",
			country: "U.K.",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "19203812038",
			cell: "192i39ii239",
		},
	}

	p1 := Person{
		firstName: "Jake",
		age:       25,
	}

	p2 := Person{
		firstName: "Jake",
		age:       25,
	}

	// p1.address.city = "NY"
	// p1.address.country = "USA"

	fmt.Println(p.firstName) // <= this is how we access directly struct's data
	fmt.Println(p1.age)      // <= this is how we access directly struct's data
	fmt.Println(p.fullName())
	fmt.Println(p.address)
	fmt.Println(p.cell)                           // <= we can access the cell data directly
	fmt.Println("Are p1 and p2 equal:", p2 == p1) // <= better as duplicate checker because the field have to be exactly the same between each other

	// Anonymous Structs
	// user := struct {
	// 	username string
	// 	email    string
	// }{
	// 	username: "user123",
	// 	email:    "example@mail.org",
	// }

	// fmt.Println(user)
	fmt.Println("Before ++", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After ++", p1.age)
}

/*
	--- Structs ---
		- Structs are defined using the 'type' and 'struct' keywords followed by curly braces '{}' containing a list of fields.
		- Fields are defined with a name and a type.
		- Anonymous Structs
		- Anonymous Fields
		- Methods :
			func (value/pointer receiver) methodName(arguments, if any...) <return tpe, if any> {
				// Method implementation
			}
		- Method Declaration
			# Value receiver method:
				func (t Type) methodName() {
					// Method implementation
				}
			# Pointer receiver method
				func (t *Type) methodName() {
					// Method implementation
				}
		- Comparing Structs
*/
