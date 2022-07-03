package main

import "fmt"

type Man struct {
	name string
}

func (man *Man) Married() {
	man.name = "Mr. " + man.name
}

func main() {
	person := Man{"panpan"}
	person.Married()

	fmt.Println(person.name)
}
