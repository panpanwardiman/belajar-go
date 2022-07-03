package main

import "fmt"

/*
type data seperti array, tapi bisa nambung berbagai jenis tipe data.
*/

type Customer struct {
	name, address string
	age           int
}

func main() {
	var person Customer
	person.name = "panpan"
	person.address = "kuningan"
	person.age = 27

	fmt.Println(person)

	person2 := Customer{
		name:    "wardiman",
		address: "karawang",
		age:     28,
	}
	fmt.Println(person2)

	john := Customer{
		name: "John",
	}
	john.sayHello()
}

func (customer Customer) sayHello() {
	fmt.Println("hello", customer.name)
}
