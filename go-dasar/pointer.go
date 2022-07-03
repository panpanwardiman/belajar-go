package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	address1 := Address{"Kuningan", "Jabar", "Indonesia"}
	address2 := &address1

	address2.city = "karawang"

	fmt.Println(address1)
	fmt.Println(address2)
}