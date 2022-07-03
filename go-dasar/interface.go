package main

import "fmt"

// Kontrak

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello ", hasName.GetName())
}

type Person struct {
	name string
}

func (person Person) GetName() string {
	return person.name
}

func main() {
	var person Person
	person.name = "panpan"

	sayHello(person)
}
