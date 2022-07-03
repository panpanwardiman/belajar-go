package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "panpan",
		"age":  "27",
	}

	fmt.Println(person["age"])

	// declare map
	book := make(map[string]string)

	book["name"] = "Buku 1"
	fmt.Println(book)

	delete(book, "name")
	fmt.Println((book))

}
