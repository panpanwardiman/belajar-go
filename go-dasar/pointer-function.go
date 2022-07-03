package main

import "fmt"

type Address struct {
	city, province, country string
}

func ChangeCountry(address *Address) {
	address.country = "indonesia"
}

func main() {
	address := Address{"kuningan", "jabar", ""}

	alamat := &address
	ChangeCountry(alamat)
	fmt.Println(alamat)
}
